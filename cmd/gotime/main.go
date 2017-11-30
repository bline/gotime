package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"

	"github.com/bline/gotime/api/proto"
	"github.com/bline/gotime/config"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/bline/gotime"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"strings"
)

func getEnv() string {
	var env = "development"
	val, ok := os.LookupEnv("GOTIME_ENV")
	if ok {
		env = val
	}
	return env
}

func getCfg(env string) *config.Config {
	cfg, err := config.New(env)
	if err != nil {
		log.Fatal("error loading config ", env, ": ", err)
	}
	return cfg
}

var (
	env = getEnv()
	cfg = getCfg(env)
)

var (
	tlsCertFilePath    = flag.String("tls_cert_file", cfg.GetString("Certs.ServerCrt"), "Path to the CRT/PEM file.")
	tlsKeyFilePath     = flag.String("tls_key_file", cfg.GetString("Certs.ServerKey"), "Path to the private key file.")
	serverHostOverride = flag.String("server_host_override", cfg.GetString("Domain"), "The server name use to verify the hostname returned by TLS handshake")
)

func serveGrpc() error {
	enableTls := cfg.GetBool("Grpc.EnableTls")

	port := cfg.GetInt("Grpc.Port")

	grpcServer := grpc.NewServer()
	api.RegisterTimeSheetServer(grpcServer, &gotime.TimeSheetService{})
	api.RegisterAccountsServer(grpcServer, &gotime.AccountsService{})
	grpclog.SetLogger(log.New(os.Stdout, "grpc: ", log.LstdFlags))

	wrappedServer := grpcweb.WrapServer(grpcServer)
	handler := func(resp http.ResponseWriter, req *http.Request) {
		wrappedServer.ServeHTTP(resp, req)
	}
	httpServer := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: http.HandlerFunc(handler),
	}
	log.Printf("Grpc Listening on Port %d", port)
	if enableTls {
		return httpServer.ListenAndServeTLS(*tlsCertFilePath, *tlsKeyFilePath)
	} else {
		return httpServer.ListenAndServe()
	}
}

func serveWeb() error {
	enableTls := cfg.GetBool("Http.EnableTls")

	port := cfg.GetInt("Http.Port")

	r := gin.Default()
	sc := cfg.Sub("Session")
	oc := sc.Sub("Options")

	secret1 := sc.GetString("secret1")
	secret2 := sc.GetString("secret2")
	var store sessions.Store
	if sc.GetBool("EnableRedis") {
		rc := sc.Sub("Redis")
		redisProto := rc.GetString("Protocol")
		redisAddr := rc.GetString("Address")
		redisPassword := rc.GetString("Password")
		var err error
		store, err = sessions.NewRedisStore(10, redisProto, redisAddr, redisPassword, []byte(secret1), []byte(secret2))
		if err != nil {
			log.Printf("Redis connection failed: %v", err)
			// Fall back to cookies
			store = sessions.NewCookieStore([]byte(secret1), []byte(secret2))
		}
	} else {
		store = sessions.NewCookieStore([]byte(secret1), []byte(secret2))
	}
	opts := sessions.Options{
		Path:     oc.GetString("Path"),
		Domain:   oc.GetString("Domain"),
		MaxAge:   oc.GetInt("MaxAge"),
		Secure:   oc.GetBool("Secure"),
		HttpOnly: oc.GetBool("HttpOnly"),
	}
	store.Options(opts)
	r.Use(sessions.Sessions("gotime-session", store))

	r.GET("/login", func(ctx *gin.Context) {

	})
	r.GET("/oauth2callback", func(ctx *gin.Context) {

	})
	r.Use(authorizeRequest())

	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: r,
	}
	strProto := "HTTP"
	if enableTls {
		strProto += "S"
	}
	log.Printf("%s Listening on Port %d", strProto, port)
	if enableTls {
		return httpServer.ListenAndServeTLS(*tlsCertFilePath, *tlsKeyFilePath)
	} else {
		return httpServer.ListenAndServe()
	}
}

var oacfg *oauth2.Config

func init() {
	googleAuthScopes := []string{"profile", "email"}
	domain := cfg.GetString("Domain")
	prefix := cfg.GetString("PathPrefix")
	redirectUrl := "https://" + domain
	if prefix != "" {
		redirectUrl += prefix
		if !strings.HasSuffix(prefix, "/") {
			redirectUrl += "/"
		}

	} else {
		redirectUrl += "/"
	}
	redirectUrl += "oauth2callback"
	oacfg = &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		RedirectURL:  redirectUrl,
		Scopes:       googleAuthScopes,
		Endpoint:     google.Endpoint,
	}
}

func authorizeRequest() gin.HandlerFunc {
	return func (ctx *gin.Context) {
		log.Printf("auth Req: %v", ctx.Request.RequestURI)
		log.Printf("new: " + "https://" + ctx.Request.Host + ctx.Request.RequestURI)
		session := sessions.Default(ctx)
		token := session.Get("token")
		if token == nil {
			rawUrl := "https://" + ctx.Request.Host + ctx.Request.RequestURI
			newUrl, err := url.Parse(rawUrl)
			if err != nil {
				log.Fatalf("Failed Parsing: %s", rawUrl)
			}
			curUrl := url.QueryEscape(newUrl.String())
			authUrl := "https://" + ctx.Request.Host + "/login?return=" + curUrl
			ctx.Redirect(http.StatusFound, authUrl)
			ctx.Abort()
		} else {
			// XXX validate token
			ctx.Next()
		}

	}
}

func main() {
	var g errgroup.Group

	flag.Parse()

	g.Go(serveGrpc)
	g.Go(serveWeb)

	if err := g.Wait(); err != nil {
		log.Fatalf("failed starting http server: %v", err)
	}
}

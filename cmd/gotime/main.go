package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"html/template"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/bline/gotime/api/proto"
	"github.com/bline/gotime/config"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/bline/gotime"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"strings"
	"path/filepath"
	"github.com/Joker/jade"
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

func getGrpcHandler() gin.HandlerFunc {
	grpcServer := grpc.NewServer()
	api.RegisterTimeSheetServer(grpcServer, &gotime.TimeSheetService{})
	api.RegisterAccountsServer(grpcServer, &gotime.AccountsService{})
	grpclog.SetLogger(log.New(os.Stdout, "grpc: ", log.LstdFlags))

	wrappedServer := grpcweb.WrapServer(grpcServer)
	handler := func(ctx *gin.Context) {
		var (
			resp = ctx.Writer
			req = ctx.Request
		)
		wrappedServer.ServeHTTP(resp, req)
	}
	return gin.HandlerFunc(handler)
}

func sessionMiddleware(engine *gin.Engine) gin.HandlerFunc {
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
	return gin.HandlerFunc(sessions.Sessions("gotime-session", store))
}

func serve() error {
	enableTls := cfg.GetBool("Http.EnableTls")

	port := cfg.GetInt("Http.Port")

	r := gin.Default()
	r.Use(sessionMiddleware(r))

	r.GET("/login", loginEndpoint)
	r.POST("/oauth2callback", oauth2callbackEndpoint)

	r.Use(authorizeRequestMiddleware())

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

func abortRequest(ctx *gin.Context) {
	r := ctx.Request
	if r.ProtoMajor == 2 && strings.HasPrefix(r.Header.Get("Content-Type"), "application/grpc") {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	} else {
		rawUrl := "https://" + r.Host + r.RequestURI
		newUrl, err := url.Parse(rawUrl)
		if err != nil {
			log.Fatalf("Failed Parsing: %s", rawUrl)
		}
		curUrl := url.QueryEscape(newUrl.String())
		authUrl := "https://" + r.Host + "/login?return=" + curUrl
		ctx.Redirect(http.StatusFound, authUrl)
		ctx.Abort()
	}
}

func authorizeRequestMiddleware() gin.HandlerFunc {
	return func (ctx *gin.Context) {
		log.Printf("auth Req: %v", ctx.Request.RequestURI)
		log.Printf("new: " + "https://" + ctx.Request.Host + ctx.Request.RequestURI)
		session := sessions.Default(ctx)
		payload := session.Get("user")
		if payload == nil {
			abortRequest(ctx)
		} else {
			ctx.Next()
		}

	}
}

func loginEndpoint(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login.pug", gin.H{})
}

func oauth2callbackEndpoint(ctx *gin.Context) {
	token := ctx.PostForm("idtoken")
	user, err := gotime.NewUserFromIDToken(token)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err})
	}
	session := sessions.Default(ctx)
	session.Set("user", user)
	ctx.JSON(http.StatusOK, user)
}

func loadHTMLString(engine *gin.Engine, name, html string) {
	// Using template instance to run New copies Delims and Funcs
	templ := template.Must(template.New(name).Funcs(engine.FuncMap).Parse(html))
	engine.SetHTMLTemplate(templ)
}

func loadPugFiles(engine *gin.Engine, files ...string) {
	for _, file := range files {
		htmlStr, err := jade.ParseFile(file)
		filename := filepath.Base(file)
		if err == nil {
			loadHTMLString(engine, filename, htmlStr)
		} else {
			log.Printf("Error parsing template %v: %v", file, err)
		}
	}
}
func loadPugGlob(engine *gin.Engine, glob string) {
	matches, err := filepath.Glob(glob)
	if err != nil {
		log.Printf("Error loading glob %v: %v", glob, err)
	}
	if len(matches) > 0 {
		loadPugFiles(engine, matches...)
	}

}

func main() {
	db, err := gotime.OpenDB()
	if err != nil {
		log.Fatalf("db connection error %v", err)
	}
	defer db.Close()
	if err := serve(); err != nil {
		log.Fatalf("failed starting http server: %v", err)
	}
}

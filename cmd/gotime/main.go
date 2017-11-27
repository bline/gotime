package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"

	"github.com/bline/gotime/api/proto"
	"github.com/bline/gotime/config"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
)

func getEnv() string {
	var env string = "development"
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
	env                = getEnv()
	cfg *config.Config = getCfg(env)
)

var (
	tlsCertFilePath    = flag.String("tls_cert_file", cfg.GetString("Certs.ServerCrt"), "Path to the CRT/PEM file.")
	tlsKeyFilePath     = flag.String("tls_key_file", cfg.GetString("Certs.ServerKey"), "Path to the private key file.")
	serverHostOverride = flag.String("server_host_override", cfg.GetString("Domain"), "The server name use to verify the hostname returned by TLS handshake")
)

func serveGrpc() error {
	enableTls := cfg.GetBool("TLSEnabled")

	port := 9090
	if enableTls {
		port = 9443
	}

	grpcServer := grpc.NewServer()
	api.RegisterTimeSheetServer(grpcServer, &timesheetService{})
	api.RegisterAccountsServer(grpcServer, &accountsService{})
	grpclog.SetLogger(log.New(os.Stdout, "grpc: ", log.LstdFlags))

	wrappedServer := grpcweb.WrapServer(grpcServer)
	handler := func(resp http.ResponseWriter, req *http.Request) {
		wrappedServer.ServeHTTP(resp, req)
	}
	httpServer := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: http.HandlerFunc(handler),
	}
	if enableTls {
		return httpServer.ListenAndServeTLS(tlsCertFilePath, tlsKeyFilePath)
	} else {
		return httpServer.ListenAndServe()
	}
}

func serveWeb() error {
	enableSsl := cfg.GetBool("SSLEnabled")

	port := 8080
	if *enableSsl {
		port = 8443
	}
	r := gin.Default()
	sc := cfg.Sub("Session")
	oc := sc.Sub("Options")
	store := sessions.NewRedosStore(10, "tcp", "127.0.0.1:6379", sc.GetString('P
	assword
	'), []byte(sc.GetString('S
	ecret1
	'), sc.GetString("Secret2")))
	opts := sessions.Options{
		Path:     oc.GetString("Path"),
		Domain:   oc.GetString("Domain"),
		MaxAge:   oc.GetInt("MaxAge"),
		Secure:   oc.GetBool("Secure"),
		HttpOnly: oc.GetBool("HttpOnly"),
	}
	store.Options(opts)
	r.Use(sessions.Sessions("gotime-auth", store))
	setupRoutes(r)

	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: http.HandlerFunc(handler),
	}
	if enableSsl {
		return httpServer.ListenAndServeTLS(tlsCertFilePath, tlsKeyFilePath)
	} else {
		return httpServer.ListenAndServe()
	}
}

func main() {
	g
	errgroup.Group

	flag.Parse()

	g.Go(serveGrpc)
	g.Go(serveWeb)

	if err := g.Wait(); err != nil {
		log.Fatalf("failed starting http server: %v", err)
	}
}

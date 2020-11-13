package main

import (
	"flag"
	"fmt"
	"github.com/inflion/inflion/dashboard/src/app/backend/args"
	"github.com/inflion/inflion/dashboard/src/app/backend/auth"
	authApi "github.com/inflion/inflion/dashboard/src/app/backend/auth/api"
	"github.com/inflion/inflion/dashboard/src/app/backend/auth/jwe"
	"github.com/inflion/inflion/dashboard/src/app/backend/client"
	clientapi "github.com/inflion/inflion/dashboard/src/app/backend/client/api"
	"github.com/inflion/inflion/dashboard/src/app/backend/handler"
	"github.com/spf13/pflag"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

var (
	argApiServerHost       = pflag.String("apiserver-host", "", "The address of the Inflion Apiserver")
	argInsecureBindAddress = pflag.IP("insecure-bind-address", net.IPv4(127, 0, 0, 1), "The IP address on which to serve the --insecure-port (set to 127.0.0.1 for all interfaces).")
	argInsecurePort        = pflag.Int("insecure-port", 9090, "The port to listen to for incoming HTTP requests.")
	argEnableInsecureLogin = pflag.Bool("enable-insecure-login", false, "When enabled, Dashboard login view will also be shown when Dashboard is not served over HTTPS. (default false)")
	argAuthenticationMode  = pflag.StringSlice("authentication-mode", []string{authApi.Basic.String()}, "Enables authentication options that will be reflected on login screen. Supported values: token, basic. "+
		"Note that basic option should only be used if apiserver has '--authorization-mode=ABAC' and '--basic-auth-file' flags set.")
)

func main() {
	log.SetOutput(os.Stdout)

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()

	initArgHolder()

	clientManager := client.NewClientManager("localhost")
	authManager := initAuthManager(clientManager)

	apiHandler, err := handler.CreateHTTPAPIHandler(clientManager, authManager)
	if err != nil {
		log.Printf("error")
	}

	http.Handle("/api/", apiHandler)

	log.Printf("Serving insecurely on HTTP port: %d", args.Holder.GetInsecurePort())
	addr := fmt.Sprintf("%s:%d", args.Holder.GetInsecureBindAddress(), args.Holder.GetInsecurePort())
	go func() { log.Fatal(http.ListenAndServe(addr, nil)) }()

	select {}
}

func initAuthManager(clientManager clientapi.ClientManager) authApi.AuthManager {
	// Init encryption key holder and token manager
	keyHolder := jwe.NewRSAKeyHolder()
	tokenManager := jwe.NewJWETokenManager(keyHolder)
	tokenTTL := time.Duration(args.Holder.GetTokenTTL())
	if tokenTTL != authApi.DefaultTokenTTL {
		tokenManager.SetTokenTTL(tokenTTL)
	}

	// Set token manager for client manager.
	clientManager.SetTokenManager(tokenManager)
	authModes := authApi.ToAuthenticationModes(args.Holder.GetAuthenticationMode())
	if len(authModes) == 0 {
		authModes.Add(authApi.Token)
	}

	// UI logic dictates this should be the inverse of the cli option
	authenticationSkippable := args.Holder.GetEnableSkipLogin()

	return auth.NewAuthManager(clientManager, tokenManager, authModes, authenticationSkippable)
}

func initArgHolder() {
	builder := args.GetHolderBuilder()
	builder.SetApiServerHost(*argApiServerHost)
	builder.SetInsecurePort(*argInsecurePort)
	builder.SetInsecureBindAddress(*argInsecureBindAddress)
	builder.SetEnableInsecureLogin(*argEnableInsecureLogin)
	builder.SetAuthenticationMode(*argAuthenticationMode)
}

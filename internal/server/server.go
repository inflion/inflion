// Copyright 2020 The Inflion Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package server

import (
	"context"
	"errors"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/inflion/inflion/api/graphql"
	"github.com/inflion/inflion/internal/hasura"
	"github.com/rs/cors"
	"github.com/unrolled/secure"
	"log"
	"net/http"
	"os"
	"runtime/debug"
)

type server struct {
	config         graphql.Config
	projectHandler projectHandler
}

func newServer(config graphql.Config, projectHandler projectHandler) server {
	return server{config: config, projectHandler: projectHandler}
}

func Run() {
	server, err := initServer()
	if err != nil {
		log.Fatal("initialize failed")
	}

	isDev := os.Getenv("RUN_ENV") != "production"

	port := "8080"
	if fromEnv := os.Getenv("PORT"); fromEnv != "" {
		port = fromEnv
	}
	log.Printf("Starting up on http://localhost:%s", port)

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(newCors().Handler)
	r.Use(extractUserId)

	sslOnly := newSecure(isDev)
	if isDev {
		r.Handle("/", playground.Handler("graphql", "/graphql"))
	} else {
		r.Use(sslOnly.Handler)
		r.Use(authMiddleware)
	}

	srv := handler.NewDefaultServer(graphql.NewExecutableSchema(server.config))
	srv.SetRecoverFunc(func(ctx context.Context, err interface{}) (userMessage error) {
		// send this panic somewhere
		log.Print(err)
		debug.PrintStack()
		return errors.New("user message on panic")
	})

	r.Handle("/graphql", srv)

	r.Route("/project", func(r chi.Router) {
		r.Post("/added", server.projectHandler.added)
		r.Post("/invite", server.projectHandler.sendInvitation)
		r.Post("/invitation/confirm", server.projectHandler.confirmInvitation)
	})

	log.Fatal(http.ListenAndServe(":"+port, r))
}

func newCors() *cors.Cors {
	return cors.New(cors.Options{
		AllowCredentials:   true,
		OptionsPassthrough: false,
		AllowedOrigins:     []string{"*"},
		AllowedMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "x-apollo-tracing"},
		ExposedHeaders:     []string{"Link"},
		MaxAge:             300, // Maximum value not ignored by any of major browsers
		Debug:              true,
	})
}

func newSecure(isDev bool) *secure.Secure {
	return secure.New(secure.Options{
		BrowserXssFilter:     true,
		ContentTypeNosniff:   true,
		FrameDeny:            true,
		HostsProxyHeaders:    []string{"X-Forwarded-Host"},
		IsDevelopment:        isDev,
		SSLProxyHeaders:      map[string]string{"X-Forwarded-Proto": "https"},
		SSLRedirect:          !isDev,
		STSIncludeSubdomains: true,
		STSPreload:           true,
		STSSeconds:           315360000,
	})
}

func extractUserId(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), hasura.HasuraUserIdKey, r.Header.Get(hasura.HasuraUserIdKey))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

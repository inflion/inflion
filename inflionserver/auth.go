package main

import (
	"context"
	"errors"
	"github.com/dgrijalva/jwt-go"
	grpcauth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"os"
)

func ensureValidToken(ctx context.Context) (context.Context, error) {
	token, err := grpcauth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return nil, status.Errorf(
			codes.Unauthenticated,
			"could not read auth token: %v",
			err,
		)
	}

	registeredToken, err := getRegisteredToken()
	if err != nil {
		return nil, err
	}

	if token != registeredToken {
		return nil, status.Errorf(codes.Unauthenticated, "invalid bearer token")
	}

	parser := new(jwt.Parser)
	parsedToken, _, err := parser.ParseUnverified(token, &jwt.StandardClaims{})
	if err != nil {
		return nil, status.Errorf(
			codes.Unauthenticated,
			"could not parsed auth token: %v",
			err,
		)
	}

	return setToken(ctx, parsedToken.Claims.(*jwt.StandardClaims)), nil
}

func getRegisteredToken() (string, error) {
	const envKey = "INFLION_CLIENT_TOKEN"
	if value, ok := os.LookupEnv(envKey); ok {
		return value, nil
	}
	return "", errors.New("could not read " + envKey)
}

func setToken(ctx context.Context, token *jwt.StandardClaims) context.Context {
	const tokenKey = "token"
	return context.WithValue(ctx, tokenKey, token)
}

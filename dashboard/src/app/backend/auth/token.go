package auth

import "github.com/inflion/inflion/dashboard/src/app/backend/auth/api"

// Implements Authenticator interface
type tokenAuthenticator struct {
	token string
}

// GetAuthInfo implements Authenticator interface. See Authenticator for more information.
func (a tokenAuthenticator) GetAuthInfo() (api.AuthInfo, error) {
	return api.AuthInfo{
		Token: a.token,
	}, nil
}

// NewTokenAuthenticator returns Authenticator based on LoginSpec.
func NewTokenAuthenticator(spec *api.LoginSpec) api.Authenticator {
	return &tokenAuthenticator{
		token: spec.Token,
	}
}

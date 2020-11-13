package auth

import "github.com/inflion/inflion/dashboard/src/app/backend/auth/api"

// Implements Authenticator interface
type basicAuthenticator struct {
	username string
	password string
}

// GetAuthInfo implements Authenticator interface. See Authenticator for more information.
func (a *basicAuthenticator) GetAuthInfo() (api.AuthInfo, error) {
	return api.AuthInfo{
		Username: a.username,
		Password: a.password,
	}, nil
}

// NewBasicAuthenticator returns Authenticator based on LoginSpec.
func NewBasicAuthenticator(spec *api.LoginSpec) api.Authenticator {
	return &basicAuthenticator{
		username: spec.Username,
		password: spec.Password,
	}
}

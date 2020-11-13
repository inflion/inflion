package auth

import (
	"github.com/inflion/inflion/dashboard/src/app/backend/auth/api"
	clientapi "github.com/inflion/inflion/dashboard/src/app/backend/client/api"
	"github.com/inflion/inflion/dashboard/src/app/backend/errors"
)

// Implements AuthManager interface
type authManager struct {
	tokenManager            api.TokenManager
	clientManager           clientapi.ClientManager
	authenticationModes     api.AuthenticationModes
	authenticationSkippable bool
}

// Login implements auth manager. See AuthManager interface for more information.
func (m authManager) Login(spec *api.LoginSpec) (*api.AuthResponse, error) {
	authenticator, err := m.getAuthenticator(spec)
	if err != nil {
		return nil, err
	}

	authInfo, err := authenticator.GetAuthInfo()
	if err != nil {
		return nil, err
	}

	err = m.healthCheck(authInfo)
	nonCriticalErrors, criticalError := errors.HandleError(err)
	if criticalError != nil || len(nonCriticalErrors) > 0 {
		return &api.AuthResponse{Errors: nonCriticalErrors}, criticalError
	}

	token, err := m.tokenManager.Generate(authInfo)
	if err != nil {
		return nil, err
	}

	return &api.AuthResponse{JWEToken: token, Errors: nonCriticalErrors}, nil
}

// Refresh implements auth manager. See AuthManager interface for more information.
func (m authManager) Refresh(jweToken string) (string, error) {
	return m.tokenManager.Refresh(jweToken)
}

func (m authManager) AuthenticationModes() []api.AuthenticationMode {
	return m.authenticationModes.Array()
}

func (m authManager) AuthenticationSkippable() bool {
	return m.authenticationSkippable
}

// Returns authenticator based on provided LoginSpec.
func (m authManager) getAuthenticator(spec *api.LoginSpec) (api.Authenticator, error) {
	if len(m.authenticationModes) == 0 {
		return nil, errors.NewInvalid("All authentication options disabled. Check --authentication-modes argument for more information.")
	}

	switch {
	case len(spec.Token) > 0 && m.authenticationModes.IsEnabled(api.Token):
		return NewTokenAuthenticator(spec), nil
	case len(spec.Username) > 0 && len(spec.Password) > 0 && m.authenticationModes.IsEnabled(api.Basic):
		return NewBasicAuthenticator(spec), nil
	}

	return nil, errors.NewInvalid("Not enough data to create authenticator.")
}

// Checks if user data extracted from provided AuthInfo structure is valid and user is correctly authenticated
// by inflion apiserver.
func (m authManager) healthCheck(authInfo api.AuthInfo) error {
	return m.clientManager.HasAccess(authInfo)
}

// NewAuthManager creates auth manager.
func NewAuthManager(clientManager clientapi.ClientManager, tokenManager api.TokenManager,
	authenticationModes api.AuthenticationModes, authenticationSkippable bool) api.AuthManager {
	return &authManager{
		tokenManager:            tokenManager,
		clientManager:           clientManager,
		authenticationModes:     authenticationModes,
		authenticationSkippable: authenticationSkippable,
	}
}

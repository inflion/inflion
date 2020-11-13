package api

import (
	"github.com/inflion/inflion/client/inflion"
	authApi "github.com/inflion/inflion/dashboard/src/app/backend/auth/api"
)

type ClientManager interface {
	CSRFKey() string
	InsecureClient() inflion.Interface
	SetTokenManager(manager authApi.TokenManager)
	HasAccess(authInfo authApi.AuthInfo) error
}

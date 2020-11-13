package auth

import (
	"github.com/emicklei/go-restful"
	authApi "github.com/inflion/inflion/dashboard/src/app/backend/auth/api"
	"github.com/inflion/inflion/dashboard/src/app/backend/errors"
	"github.com/inflion/inflion/dashboard/src/app/backend/validation"
	"net/http"
)

type AuthHandler struct {
	manager authApi.AuthManager
}

// Install creates new endpoints for dashboard auth, such as login. It allows user to log in to dashboard using
// one of the supported methods. See AuthManager and Authenticator for more information.
func (h AuthHandler) Install(ws *restful.WebService) {
	ws.Route(
		ws.POST("/login").
			To(h.handleLogin).
			Reads(authApi.LoginSpec{}).
			Writes(authApi.AuthResponse{}))
	ws.Route(
		ws.GET("/login/status").
			To(h.handleLoginStatus).
			Writes(validation.LoginStatus{}))
	ws.Route(
		ws.POST("/token/refresh").
			Reads(authApi.TokenRefreshSpec{}).
			To(h.handleJWETokenRefresh).
			Writes(authApi.AuthResponse{}))
}

func (h AuthHandler) handleLogin(request *restful.Request, response *restful.Response) {
	loginSpec := new(authApi.LoginSpec)
	if err := request.ReadEntity(loginSpec); err != nil {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(errors.HandleHTTPError(err), err.Error()+"\n")
		return
	}

	loginResponse, err := h.manager.Login(loginSpec)
	if err != nil {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(errors.HandleHTTPError(err), err.Error()+"\n")
		return
	}

	response.WriteHeaderAndEntity(http.StatusOK, loginResponse)
}

func (h *AuthHandler) handleLoginStatus(request *restful.Request, response *restful.Response) {
	response.WriteHeaderAndEntity(http.StatusOK, validation.ValidateLoginStatus(request))
}

func (h *AuthHandler) handleJWETokenRefresh(request *restful.Request, response *restful.Response) {
	tokenRefreshSpec := new(authApi.TokenRefreshSpec)
	if err := request.ReadEntity(tokenRefreshSpec); err != nil {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(errors.HandleHTTPError(err), err.Error()+"\n")
		return
	}

	refreshedJWEToken, err := h.manager.Refresh(tokenRefreshSpec.JWEToken)
	if err != nil {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(errors.HandleHTTPError(err), err.Error()+"\n")
		return
	}

	response.WriteHeaderAndEntity(http.StatusOK, &authApi.AuthResponse{
		JWEToken: refreshedJWEToken,
		Errors:   make([]error, 0),
	})
}

// NewAuthHandler created AuthHandler instance.
func NewAuthHandler(manager authApi.AuthManager) AuthHandler {
	return AuthHandler{manager: manager}
}

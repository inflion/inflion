package handler

import (
	"github.com/emicklei/go-restful"
	"github.com/inflion/inflion/dashboard/src/app/backend/api"
	"github.com/inflion/inflion/dashboard/src/app/backend/auth"
	authApi "github.com/inflion/inflion/dashboard/src/app/backend/auth/api"
	clientapi "github.com/inflion/inflion/dashboard/src/app/backend/client/api"
	"golang.org/x/net/xsrftoken"
	"net/http"
)

type APIHandler struct {
	cManager clientapi.ClientManager
}

func CreateHTTPAPIHandler(clientManager clientapi.ClientManager, authManager authApi.AuthManager) (http.Handler, error) {
	apiHandler := APIHandler{cManager: clientManager}

	wsContainer := restful.NewContainer()
	wsContainer.EnableContentEncoding(false)

	apiV1Ws := new(restful.WebService)
	apiV1Ws.Path("/api/v1").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)
	wsContainer.Add(apiV1Ws)

	authHandler := auth.NewAuthHandler(authManager)
	authHandler.Install(apiV1Ws)

	apiV1Ws.Route(
		apiV1Ws.GET("csrftoken/{action}").
			To(apiHandler.handleGetCsrfToken).
			Writes(api.CsrfToken{}))

	return wsContainer, nil
}

func (apiHandler *APIHandler) handleGetCsrfToken(request *restful.Request, response *restful.Response) {
	action := request.PathParameter("action")
	token := xsrftoken.Generate(apiHandler.cManager.CSRFKey(), "none", action)
	response.WriteHeaderAndEntity(http.StatusOK, api.CsrfToken{Token: token})
}
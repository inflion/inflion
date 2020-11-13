package validation

import (
	"github.com/emicklei/go-restful"
	"github.com/inflion/inflion/dashboard/src/app/backend/args"
	"github.com/inflion/inflion/dashboard/src/app/backend/client"
)

// LoginStatus is returned as a response to login status check. Used by the frontend to determine if is logged in
// and if login page should be shown.
type LoginStatus struct {
	// True when token header indicating logged in user is found in request.
	TokenPresent bool `json:"tokenPresent"`

	// True when authorization header indicating logged in user is found in request.
	HeaderPresent bool `json:"headerPresent"`

	// True if dashboard is configured to use HTTPS connection. It is required for secure
	// data exchange during login operation.
	HTTPSMode bool `json:"httpsMode"`

	// True if impersonation is enabled
	ImpersonationPresent bool `json:"impersonationPresent"`

	// The impersonated user
	ImpersonatedUser string `json:"impersonatedUser"`
}

// ValidateLoginStatus returns information about user login status and if request was made over HTTPS.
func ValidateLoginStatus(request *restful.Request) *LoginStatus {
	authHeader := request.HeaderParameter("Authorization")
	tokenHeader := request.HeaderParameter(client.JWETokenHeader)
	impersonationHeader := request.HeaderParameter("Impersonate-User")

	httpsMode := request.Request.TLS != nil
	if args.Holder.GetEnableInsecureLogin() {
		httpsMode = true
	}

	loginStatus := &LoginStatus{
		TokenPresent:         len(tokenHeader) > 0,
		HeaderPresent:        len(authHeader) > 0,
		ImpersonationPresent: len(impersonationHeader) > 0,
		HTTPSMode:            httpsMode,
	}

	if loginStatus.ImpersonationPresent {
		loginStatus.ImpersonatedUser = impersonationHeader
	}

	return loginStatus
}

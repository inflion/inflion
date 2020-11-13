package api

// CsrfToken is used to secure requests from CSRF attacks
type CsrfToken struct {
	// Token generated on request for validation
	Token string `json:"token"`
}
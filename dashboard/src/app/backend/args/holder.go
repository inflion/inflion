package args

import (
	"net"
)

var Holder = &holder{}

type holder struct {
	insecurePort        int
	insecureBindAddress net.IP
	apiServerHost       string
	enableInsecureLogin bool
	tokenTTL            int
	authenticationMode  []string
	enableSkipLogin     bool
}

func (h *holder) GetInsecurePort() int {
	return h.insecurePort
}

func (h *holder) GetInsecureBindAddress() net.IP {
	return h.insecureBindAddress
}

func (h *holder) GetApiServerHost() string {
	return h.apiServerHost
}

func (h *holder) GetEnableInsecureLogin() bool {
	return h.enableInsecureLogin
}

func (h *holder) GetTokenTTL() int {
	return h.tokenTTL
}

func (h *holder) GetAuthenticationMode() []string {
	return h.authenticationMode
}

func (h *holder) GetEnableSkipLogin() bool {
	return h.enableSkipLogin
}

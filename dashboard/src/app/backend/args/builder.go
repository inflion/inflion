package args

import "net"

var builder = &holderBuilder{holder: Holder}

type holderBuilder struct {
	holder *holder
}

func (h *holderBuilder) SetApiServerHost(apiServerHost string) *holderBuilder {
	h.holder.apiServerHost = apiServerHost
	return h
}

func (h *holderBuilder) SetInsecurePort(port int) *holderBuilder {
	h.holder.insecurePort = port
	return h
}

func (h *holderBuilder) SetInsecureBindAddress(ip net.IP) *holderBuilder {
	h.holder.insecureBindAddress = ip
	return h
}

func (h *holderBuilder) SetEnableInsecureLogin(enableInsecureLogin bool) *holderBuilder {
	h.holder.enableInsecureLogin = enableInsecureLogin
	return h
}

func (h *holderBuilder) SetAuthenticationMode(authenticationMode []string) *holderBuilder {
	h.holder.authenticationMode = authenticationMode
	return h
}

// GetHolderBuilder returns singleton instance of argument holder builder.
func GetHolderBuilder() *holderBuilder {
	return builder
}

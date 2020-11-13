package api

// ToAuthenticationModes transforms array of authentication mode strings to valid AuthenticationModes type.
func ToAuthenticationModes(modes []string) AuthenticationModes {
	result := AuthenticationModes{}
	modesMap := map[string]bool{}

	for _, mode := range []AuthenticationMode{Token, Basic} {
		modesMap[mode.String()] = true
	}

	for _, mode := range modes {
		if _, exists := modesMap[mode]; exists {
			result.Add(AuthenticationMode(mode))
		}
	}

	return result
}


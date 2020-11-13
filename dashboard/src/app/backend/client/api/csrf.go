package api

import "crypto/rand"

// GenerateCSRFKey generates random csrf key
func GenerateCSRFKey() string {
	bytes := make([]byte, 256)
	_, err := rand.Read(bytes)
	if err != nil {
		panic("could not generate csrf key")
	}

	return string(bytes)
}

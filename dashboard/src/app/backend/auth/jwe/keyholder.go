package jwe

import (
	"crypto/rand"
	"crypto/rsa"
	"gopkg.in/square/go-jose.v2"
	"log"
	"sync"
)

// KeyHolder is responsible for generating, storing and synchronizing encryption key used for token
// generation/decryption.
type KeyHolder interface {
	// Returns encrypter instance that can be used to encrypt data.
	Encrypter() jose.Encrypter

	// Returns encryption key that can be used to decrypt data.
	Key() *rsa.PrivateKey

	// Forces refresh of encryption key synchronized with kubernetes resource (secret).
	Refresh()
}

// Implements KeyHolder interface
type rsaKeyHolder struct {
	// 256-byte random RSA key pair. Synced with a key saved in a secret.
	key *rsa.PrivateKey
	mux sync.Mutex
}

// Encrypter implements key holder interface. See KeyHolder for more information.
// Used encryption algorithms:
//    - Content encryption: AES-GCM (256)
//    - Key management: RSA-OAEP-SHA256
func (h *rsaKeyHolder) Encrypter() jose.Encrypter {
	publicKey := &h.Key().PublicKey
	encrypter, err := jose.NewEncrypter(jose.A256GCM, jose.Recipient{Algorithm: jose.RSA_OAEP_256, Key: publicKey}, nil)
	if err != nil {
		panic(err)
	}

	return encrypter
}

// Key implements key holder interface. See KeyHolder for more information.
func (h *rsaKeyHolder) Key() *rsa.PrivateKey {
	h.mux.Lock()
	defer h.mux.Unlock()
	return h.key
}

// Refresh implements key holder interface. See KeyHolder for more information.
func (h *rsaKeyHolder) Refresh() {
	// TODO implements for future release.
}

func (h *rsaKeyHolder) init() {
	h.initEncryptionKey()
}

// Generates encryption key used to encrypt token payload.
func (h *rsaKeyHolder) initEncryptionKey() {
	log.Print("Generating JWE encryption key")
	h.mux.Lock()
	defer h.mux.Unlock()

	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}

	h.key = privateKey
}

// NewRSAKeyHolder creates new KeyHolder instance.
func NewRSAKeyHolder() KeyHolder {
	holder := &rsaKeyHolder{}
	holder.init()
	return holder
}

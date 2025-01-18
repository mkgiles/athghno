package streams

type PublicKey struct {
	Registrable  `json:"-"`
	Id           string `json:"id"`
	Type         string `json:"type"`
	Owner        string `json:"owner"`
	PublicKeyPem string `json:"publicKeyPem"`
}

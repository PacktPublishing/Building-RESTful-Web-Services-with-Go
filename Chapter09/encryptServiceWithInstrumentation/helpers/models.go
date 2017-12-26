package helpers

// EncryptRequest strctures request coming from client
type EncryptRequest struct {
	Text string `json:"text"`
	Key  string `json:"key"`
}

// EncryptResponse strctures response going to the client
type EncryptResponse struct {
	Message string `json:"message"`
	Err     string `json:"error"`
}

// DecryptRequest strctures request coming from client
type DecryptRequest struct {
	Message string `json:"message"`
	Key     string `json:"key"`
}

// DecryptResponse strctures response going to the client
type DecryptResponse struct {
	Text string `json:"text"`
	Err  string `json:"error"`
}

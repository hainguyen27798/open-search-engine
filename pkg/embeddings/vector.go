package embeddings

type Vector struct {
	Embeddings [][]float32 `json:"embeddings"`
	Model      string      `json:"model"`
	Usage      interface{} `json:"usage"`
	Error      string      `json:"error"`
	Detail     string      `json:"detail"`
}

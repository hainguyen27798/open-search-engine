package embeddings

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/hainguyen27798/open-search-engine/global"
	"io"
	"log"
	"net/http"
)

func TextToVector(input string) (*[][]float32, error) {
	body := map[string]interface{}{
		"model": global.Config.EmbeddingModel.TextEMName,
		"input": input,
	}
	jsonBytes, err := json.Marshal(body)
	if err != nil {
		log.Printf(err.Error())
		return nil, err
	}
	resp, err := http.Post(
		global.Config.EmbeddingModel.TextEMURL,
		"application/json",
		bytes.NewBuffer(jsonBytes),
	)
	if err != nil {
		log.Printf(err.Error())
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	data, _ := io.ReadAll(resp.Body)
	var vectorData Vector
	if err := json.Unmarshal(data, &vectorData); err != nil {
		log.Printf(err.Error())
		return nil, err
	}

	if vectorData.Error != "" {
		return nil, errors.New(vectorData.Error)
	}

	return &vectorData.Embeddings, nil
}

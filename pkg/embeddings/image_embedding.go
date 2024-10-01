package embeddings

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/hainguyen27798/open-search-engine/global"
	"io"
	"log"
	"mime/multipart"
	"net/http"
)

func ImageToVector(url string) (*[][]float32, error) {
	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	_ = writer.WriteField("model", global.Config.EmbeddingModel.ImageEMName)
	_ = writer.WriteField("urls", url)

	_ = writer.Close()

	req, err := http.NewRequest(
		"POST",
		global.Config.EmbeddingModel.ImageEMURL,
		&buf,
	)
	if err != nil {
		log.Printf(err.Error())
		return nil, err
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", global.Config.EmbeddingModel.Auth)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf(err.Error())
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	body, _ := io.ReadAll(resp.Body)

	var vectorData Vector
	if err := json.Unmarshal(body, &vectorData); err != nil {
		log.Printf(err.Error())
		return nil, err
	}

	if vectorData.Error != "" {
		return nil, errors.New(vectorData.Error)
	}

	if vectorData.Detail != "" {
		return nil, errors.New(vectorData.Detail)
	}

	return &vectorData.Embeddings, nil
}

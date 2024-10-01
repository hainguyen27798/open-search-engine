package tests

import (
	"github.com/hainguyen27798/open-search-engine/internal/initialize"
	"github.com/hainguyen27798/open-search-engine/pkg/embeddings"
	"testing"
)

func TestConvertImageToVector(t *testing.T) {
	initialize.LoadEnv()

	vectorData, err := embeddings.ImageToVector("https://picsum.photos/200/300")
	if err != nil {
		t.Fatalf(`Error mesage: %v`, err)
	}

	t.Logf("Result: %v", *vectorData)
}

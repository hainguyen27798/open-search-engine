package tests

import (
	"github.com/hainguyen27798/open-search-engine/internal/infrastructure/initialize"
	"github.com/hainguyen27798/open-search-engine/pkg/embeddings"
	"testing"
)

func TestConvertTextToVector(t *testing.T) {
	initialize.LoadEnv()

	vectorData, err := embeddings.TextToVector("test")
	if err != nil {
		t.Fatalf(`Error mesage: %v`, err)
	}

	t.Logf("Result: %v", *vectorData)
}

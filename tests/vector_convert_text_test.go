package tests

import (
	"github.com/hainguyen27798/open-search-engine/internal/infrastructure/initialize"
	"github.com/hainguyen27798/open-search-engine/pkg/embeddings"
	"testing"
)

func TestConvertTextToVector(t *testing.T) {
	initialize.LoadEnv()

	vectorData, err := embeddings.TextToVector("test demo")
	if err != nil {
		t.Fatalf(`Error mesage: %v`, err)
	}

	rs := *vectorData
	t.Logf("Result: %v", len(rs))
	if len(rs) >= 1 {
		t.Logf("Dimensional Size: %v", len(rs[0]))
	}
}

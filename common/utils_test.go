package common

import (
	"testing"

	"github.com/mark3labs/mcp-go/mcp"
)

// JSON numbers decode into interface{} as float64, so numeric params such as
// {"page": 1} previously panicked ExtractParams via an unchecked .(string).
func TestExtractParamsCoercesNonStringValues(t *testing.T) {
	req := mcp.CallToolRequest{}
	req.Params.Arguments = map[string]any{
		"params": map[string]interface{}{
			"page":      float64(1),
			"page_size": float64(50),
			"direction": "all",
			"active":    true,
		},
	}

	got := ExtractParams(req, []string{"page", "page_size", "direction", "active"})
	want := map[string]string{
		"page":      "1",
		"page_size": "50",
		"direction": "all",
		"active":    "true",
	}
	for k, v := range want {
		if got[k] != v {
			t.Errorf("param %q = %q, want %q", k, got[k], v)
		}
	}
}

func TestExtractParamsMissingParamsKey(t *testing.T) {
	req := mcp.CallToolRequest{}
	req.Params.Arguments = map[string]any{}
	if got := ExtractParams(req, []string{"page"}); len(got) != 0 {
		t.Errorf("expected empty map, got %v", got)
	}
}

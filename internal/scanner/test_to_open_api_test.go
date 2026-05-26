//ff:func feature=scan type=convert control=sequence
//ff:what TestToOpenAPI 테스트
package scanner

import (
	"strings"
	"testing"
)

func TestToOpenAPI(t *testing.T) {
	result := &ScanResult{
		Endpoints: []Endpoint{
			{
				Method:  "GET",
				Path:    "/api/v1/users",
				Handler: "handler.go:h.ListUsers",
			},
		},
	}

	out, err := ToOpenAPI(result, nil)
	if err != nil {
		t.Fatalf("ToOpenAPI() error: %v", err)
	}
	if !strings.Contains(string(out), "openapi") {
		t.Error("expected 'openapi' in output")
	}
}

//ff:func feature=scan type=convert control=sequence
//ff:what TestToOpenAPI_Empty 테스트
package scanner

import (
	"strings"
	"testing"
)

func TestToOpenAPI_Empty(t *testing.T) {
	result := &ScanResult{}
	out, err := ToOpenAPI(result, nil)
	if err != nil {
		t.Fatalf("ToOpenAPI() error: %v", err)
	}
	if !strings.Contains(string(out), "openapi") {
		t.Error("expected 'openapi' in output")
	}
}

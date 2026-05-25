//ff:func feature=scan type=extract control=sequence
//ff:what TestBuildResponses_WithResponses 테스트
package scanner

import "testing"

func TestBuildResponses_WithResponses(t *testing.T) {
	resps := []Response{{Status: "200"}, {Status: "400"}}
	schemas := map[string]any{}
	result := buildResponses(resps, schemas)
	if result == nil {
		t.Fatal("expected non-nil")
	}
}

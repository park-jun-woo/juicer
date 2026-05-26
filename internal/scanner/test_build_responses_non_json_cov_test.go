//ff:func feature=scan type=test control=sequence
//ff:what TestBuildResponses_NonJSONCov 테스트
package scanner

import "testing"

func TestBuildResponses_NonJSONCov(t *testing.T) {
	resps := []Response{{Status: "204", Kind: "empty"}}
	schemas := map[string]any{}
	result := buildResponses(resps, schemas)
	if result["204"] == nil {
		t.Fatal("expected 204 response")
	}
}

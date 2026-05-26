//ff:func feature=scan type=test control=sequence
//ff:what TestBuildResponses_Empty 테스트
package scanner

import "testing"

func TestBuildResponses_Empty(t *testing.T) {
	schemas := map[string]any{}
	result := buildResponses(nil, schemas)
	if result == nil {
		t.Fatal("expected non-nil")
	}
}

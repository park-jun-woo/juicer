//ff:func feature=scan type=test control=sequence
//ff:what TestExtractPathString_Concat 테스트
package fiber

import "testing"

func TestExtractPathString_Concat(t *testing.T) {
	got, ok := extractPathFor(t, `"/api" + "/v1"`)
	if !ok || got != "/api/v1" {
		t.Fatalf("got %q ok=%v", got, ok)
	}
}

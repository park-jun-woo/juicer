//ff:func feature=scan type=test control=sequence
//ff:what TestExtractPathString_Literal 테스트
package fiber

import "testing"

func TestExtractPathString_Literal(t *testing.T) {
	got, ok := extractPathFor(t, `"/users"`)
	if !ok || got != "/users" {
		t.Fatalf("got %q ok=%v", got, ok)
	}
}

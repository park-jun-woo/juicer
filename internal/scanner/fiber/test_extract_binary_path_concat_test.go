//ff:func feature=scan type=test control=sequence
//ff:what TestExtractBinaryPath_Concat 테스트
package fiber

import "testing"

func TestExtractBinaryPath_Concat(t *testing.T) {
	got, ok := extractBinaryPath(binExpr(t, `"/api" + "/users"`))
	if !ok || got != "/api/users" {
		t.Fatalf("got %q ok=%v", got, ok)
	}
}

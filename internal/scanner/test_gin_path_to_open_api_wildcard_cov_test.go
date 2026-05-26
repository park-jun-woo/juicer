//ff:func feature=scan type=test control=sequence
//ff:what TestGinPathToOpenAPI_WildcardCov 테스트
package scanner

import "testing"

func TestGinPathToOpenAPI_WildcardCov(t *testing.T) {
	got := ginPathToOpenAPI("/api/files/*filepath")
	if got != "/api/files/{filepath}" {
		t.Fatalf("expected /api/files/{filepath}, got %s", got)
	}
}

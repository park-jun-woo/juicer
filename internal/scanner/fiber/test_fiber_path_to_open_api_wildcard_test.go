//ff:func feature=scan type=test control=sequence
//ff:what TestFiberPathToOpenAPI_Wildcard 테스트
package fiber

import "testing"

func TestFiberPathToOpenAPI_Wildcard(t *testing.T) {
	got := fiberPathToOpenAPI("/static/*filepath")
	if got != "/static/{filepath}" {
		t.Fatalf("expected /static/{filepath}, got %s", got)
	}
}

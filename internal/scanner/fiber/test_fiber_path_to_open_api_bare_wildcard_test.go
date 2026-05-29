//ff:func feature=scan type=test control=sequence
//ff:what TestFiberPathToOpenAPI_BareWildcard 테스트
package fiber

import "testing"

func TestFiberPathToOpenAPI_BareWildcard(t *testing.T) {
	got := fiberPathToOpenAPI("/*")
	if got != "/{wildcard}" {
		t.Fatalf("expected /{wildcard}, got %s", got)
	}
}

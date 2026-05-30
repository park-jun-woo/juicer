//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestAssignSubAppPrefix_SameFile 테스트
package hono

import "testing"

func TestAssignSubAppPrefix_SameFile(t *testing.T) {
	prefixMap := map[string]string{}
	honoVars := map[string]map[string]bool{}

	if !assignSubAppPrefix("a.ts", "a.ts", "users", "/api/users", prefixMap, honoVars) {
		t.Fatal("expected change")
	}

	if assignSubAppPrefix("a.ts", "a.ts", "users", "/api/users", prefixMap, honoVars) {
		t.Fatal("expected no change on identical value")
	}
}

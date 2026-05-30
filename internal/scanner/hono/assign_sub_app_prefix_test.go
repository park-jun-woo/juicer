//ff:func feature=scan type=test control=sequence topic=hono
//ff:what assignSubAppPrefix 테스트
package hono

import "testing"

func TestAssignSubAppPrefix_SameFile(t *testing.T) {
	prefixMap := map[string]string{}
	honoVars := map[string]map[string]bool{}

	// new entry -> changed
	if !assignSubAppPrefix("a.ts", "a.ts", "users", "/api/users", prefixMap, honoVars) {
		t.Fatal("expected change")
	}
	// same value -> no change
	if assignSubAppPrefix("a.ts", "a.ts", "users", "/api/users", prefixMap, honoVars) {
		t.Fatal("expected no change on identical value")
	}
}

func TestAssignSubAppPrefix_CrossFile(t *testing.T) {
	prefixMap := map[string]string{}
	honoVars := map[string]map[string]bool{"def.ts": {"r1": true, "r2": true}}

	if !assignSubAppPrefix("def.ts", "src.ts", "sub", "/v1", prefixMap, honoVars) {
		t.Fatal("expected cross-file change")
	}
	if prefixMap[prefixKey("def.ts", "r1")] != "/v1" || prefixMap[prefixKey("def.ts", "r2")] != "/v1" {
		t.Fatalf("cross-file vars not all set: %v", prefixMap)
	}
	// second apply with same value -> no change
	if assignSubAppPrefix("def.ts", "src.ts", "sub", "/v1", prefixMap, honoVars) {
		t.Fatal("expected no change on identical cross-file value")
	}
}

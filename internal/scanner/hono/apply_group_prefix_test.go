//ff:func feature=scan type=test control=sequence topic=hono
//ff:what applyGroupPrefix 테스트
package hono

import "testing"

func TestApplyGroupPrefix_SameFile(t *testing.T) {
	g := routeGroup{
		Prefix:     "/users",
		ParentVar:  "app",
		SubAppName: "users",
		SourceFile: "app.ts",
	}
	prefixMap := map[string]string{prefixKey("app.ts", "app"): "/api"}
	honoVars := map[string]map[string]bool{}
	imports := map[string]map[string]string{}

	changed := applyGroupPrefix(g, prefixMap, honoVars, imports)
	if !changed {
		t.Fatal("expected change")
	}
	if prefixMap[prefixKey("app.ts", "users")] != "/api/users" {
		t.Fatalf("sub-app prefix = %q", prefixMap[prefixKey("app.ts", "users")])
	}

	// applying again -> no change
	if applyGroupPrefix(g, prefixMap, honoVars, imports) {
		t.Fatal("expected no change on second apply")
	}
}

func TestApplyGroupPrefix_CrossFile(t *testing.T) {
	g := routeGroup{
		Prefix:     "/v1",
		ParentVar:  "app",
		SubAppName: "sub",
		SourceFile: "app.ts",
	}
	prefixMap := map[string]string{}
	// sub is imported from routes.ts
	imports := map[string]map[string]string{"app.ts": {"sub": "routes.ts"}}
	// routes.ts declares a hono var "r"
	honoVars := map[string]map[string]bool{"routes.ts": {"r": true}}

	changed := applyGroupPrefix(g, prefixMap, honoVars, imports)
	if !changed {
		t.Fatal("expected cross-file change")
	}
	if prefixMap[prefixKey("routes.ts", "r")] != "/v1" {
		t.Fatalf("cross-file prefix = %q", prefixMap[prefixKey("routes.ts", "r")])
	}
}

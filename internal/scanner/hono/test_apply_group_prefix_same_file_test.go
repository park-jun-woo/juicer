//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestApplyGroupPrefix_SameFile 테스트
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

	if applyGroupPrefix(g, prefixMap, honoVars, imports) {
		t.Fatal("expected no change on second apply")
	}
}

//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestApplyGroupPrefix_CrossFile 테스트
package hono

import "testing"

func TestApplyGroupPrefix_CrossFile(t *testing.T) {
	g := routeGroup{
		Prefix:     "/v1",
		ParentVar:  "app",
		SubAppName: "sub",
		SourceFile: "app.ts",
	}
	prefixMap := map[string]string{}

	imports := map[string]map[string]string{"app.ts": {"sub": "routes.ts"}}

	honoVars := map[string]map[string]bool{"routes.ts": {"r": true}}

	changed := applyGroupPrefix(g, prefixMap, honoVars, imports)
	if !changed {
		t.Fatal("expected cross-file change")
	}
	if prefixMap[prefixKey("routes.ts", "r")] != "/v1" {
		t.Fatalf("cross-file prefix = %q", prefixMap[prefixKey("routes.ts", "r")])
	}
}

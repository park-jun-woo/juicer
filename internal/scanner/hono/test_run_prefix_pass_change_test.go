//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestRunPrefixPass_Change 테스트
package hono

import "testing"

func TestRunPrefixPass_Change(t *testing.T) {

	groups := []routeGroup{{
		Prefix:     "/sub",
		ParentVar:  "app",
		SubAppName: "sub",
		SourceFile: "f.ts",
	}}
	prefixMap := map[string]string{
		prefixKey("f.ts", "app"): "/api",
	}
	honoVars := map[string]map[string]bool{
		"f.ts": {"app": true, "sub": true},
	}
	imports := map[string]map[string]string{}
	changed := runPrefixPass(groups, prefixMap, honoVars, imports)
	if !changed {
		t.Fatal("expected change on first pass")
	}

	if runPrefixPass(groups, prefixMap, honoVars, imports) {
		t.Fatal("expected stable on second pass")
	}
}

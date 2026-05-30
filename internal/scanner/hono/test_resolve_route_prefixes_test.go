//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestResolveRoutePrefixes 테스트
package hono

import "testing"

func TestResolveRoutePrefixes(t *testing.T) {
	groups := []routeGroup{{
		Prefix:     "/sub",
		ParentVar:  "app",
		SubAppName: "sub",
		SourceFile: "f.ts",
	}}
	basePaths := map[string]string{
		prefixKey("f.ts", "app"): "/api",
	}
	honoVars := map[string]map[string]bool{
		"f.ts": {"app": true, "sub": true},
	}
	pm := resolveRoutePrefixes(groups, basePaths, honoVars, map[string]map[string]string{})
	if pm[prefixKey("f.ts", "app")] != "/api" {
		t.Fatalf("base prefix lost: %v", pm)
	}
	if pm[prefixKey("f.ts", "sub")] == "" {
		t.Fatalf("sub prefix not propagated: %v", pm)
	}
}

//ff:func feature=scan type=test control=sequence topic=hono
//ff:what runPrefixPass / resolveRoutePrefixes 테스트
package hono

import "testing"

func TestRunPrefixPass_NoChange(t *testing.T) {
	// no groups -> no change
	prefixMap := map[string]string{}
	if runPrefixPass(nil, prefixMap, map[string]map[string]bool{}, map[string]map[string]string{}) {
		t.Fatal("expected no change for empty groups")
	}
}

func TestRunPrefixPass_Change(t *testing.T) {
	// parent app has basePath; group maps sub-app under prefix
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
	// second pass should be stable (no further change)
	if runPrefixPass(groups, prefixMap, honoVars, imports) {
		t.Fatal("expected stable on second pass")
	}
}

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

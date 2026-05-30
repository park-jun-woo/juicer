//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestBuildSchemaSrcMap_Empty 테스트
package hono

import "testing"

func TestBuildSchemaSrcMap_Empty(t *testing.T) {
	ctx := minimalCtx()
	if m := buildSchemaSrcMap(ctx); len(m) != 0 {
		t.Fatalf("expected empty map, got %d", len(m))
	}
}

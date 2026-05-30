//ff:func feature=scan type=test control=sequence topic=hono
//ff:what buildSchemaSrcMap 테스트
package hono

import "testing"

func TestBuildSchemaSrcMap(t *testing.T) {
	fi := mustParse(t, []byte(`const userSchema = z.object({ name: z.string() });`+"\n"))
	ctx := minimalCtx()
	ctx.parsed["a.ts"] = fi

	srcMap := buildSchemaSrcMap(ctx)
	if _, ok := srcMap["userSchema"]; !ok {
		t.Fatalf("expected userSchema in src map, got %v", keysOf(srcMap))
	}
}

func TestBuildSchemaSrcMap_Empty(t *testing.T) {
	ctx := minimalCtx()
	if m := buildSchemaSrcMap(ctx); len(m) != 0 {
		t.Fatalf("expected empty map, got %d", len(m))
	}
}

func keysOf(m map[string][]byte) []string {
	var k []string
	for key := range m {
		k = append(k, key)
	}
	return k
}

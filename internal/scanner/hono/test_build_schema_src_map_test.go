//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestBuildSchemaSrcMap 테스트
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

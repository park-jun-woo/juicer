//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestMergeFileSchemas_Merges 테스트
package express

import "testing"

func TestMergeFileSchemas_Merges(t *testing.T) {
	fi := mustParse(t, []byte(`const userSchema = z.object({ name: z.string() });`))
	ctx := newSchemaCtx()
	mergeFileSchemas(ctx, fi)
	if _, ok := ctx.schemas["userSchema"]; !ok {
		t.Fatalf("schema not merged: %v", ctx.schemas)
	}
	if _, ok := ctx.schemaSrc["userSchema"]; !ok {
		t.Fatalf("schemaSrc not merged")
	}
}

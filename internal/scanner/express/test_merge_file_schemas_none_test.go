//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestMergeFileSchemas_None 테스트
package express

import "testing"

func TestMergeFileSchemas_None(t *testing.T) {
	fi := mustParse(t, []byte(`const x = 1;`))
	ctx := newSchemaCtx()
	mergeFileSchemas(ctx, fi)
	if len(ctx.schemas) != 0 {
		t.Fatalf("expected no schemas, got %v", ctx.schemas)
	}
}

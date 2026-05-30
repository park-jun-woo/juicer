//ff:func feature=scan type=test control=sequence topic=express
//ff:what mergeFileSchemas: 스키마 병합 / 스키마 없음
package express

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
)

func newSchemaCtx() *scanContext {
	return &scanContext{
		schemas:   map[string]*sitter.Node{},
		schemaSrc: map[string][]byte{},
	}
}

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

func TestMergeFileSchemas_None(t *testing.T) {
	fi := mustParse(t, []byte(`const x = 1;`))
	ctx := newSchemaCtx()
	mergeFileSchemas(ctx, fi)
	if len(ctx.schemas) != 0 {
		t.Fatalf("expected no schemas, got %v", ctx.schemas)
	}
}

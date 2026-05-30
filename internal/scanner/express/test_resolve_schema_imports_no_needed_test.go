//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestResolveSchemaImports_NoNeeded 테스트
package express

import "testing"

func TestResolveSchemaImports_NoNeeded(t *testing.T) {
	ctx := newScanCtx(t.TempDir())

	resolveSchemaImports(ctx)
	if len(ctx.schemas) != 0 {
		t.Fatalf("expected no schemas, got %v", ctx.schemas)
	}
}

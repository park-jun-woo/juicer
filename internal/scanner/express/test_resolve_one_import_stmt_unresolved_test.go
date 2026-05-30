//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestResolveOneImportStmt_Unresolved 테스트
package express

import (
	"path/filepath"
	"testing"
)

func TestResolveOneImportStmt_Unresolved(t *testing.T) {
	dir := t.TempDir()
	fi := mustParse(t, []byte(`import { userSchema } from 'external';`))
	fi.Path = filepath.Join(dir, "app.ts")
	ctx := newScanCtx(dir)
	resolveOneImportStmt(ctx, fi, firstImportStmt(t, fi), map[string]bool{"userSchema": true})
	if len(ctx.schemas) != 0 {
		t.Fatalf("expected no merge, got %v", ctx.schemas)
	}
}

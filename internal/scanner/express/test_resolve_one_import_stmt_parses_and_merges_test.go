//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestResolveOneImportStmt_ParsesAndMerges 테스트
package express

import (
	"path/filepath"
	"testing"
)

func TestResolveOneImportStmt_ParsesAndMerges(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "schemas.ts", `export const userSchema = z.object({ name: z.string() });`)
	fi := mustParse(t, []byte(`import { userSchema } from './schemas';`))
	fi.Path = filepath.Join(dir, "app.ts")
	ctx := newScanCtx(dir)
	resolveOneImportStmt(ctx, fi, firstImportStmt(t, fi), map[string]bool{"userSchema": true})
	if _, ok := ctx.schemas["userSchema"]; !ok {
		t.Fatalf("schema not merged: %v", ctx.schemas)
	}
}

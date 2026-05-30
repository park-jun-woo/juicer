//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestResolveOneImportStmt_AlreadyParsed 테스트
package express

import (
	"path/filepath"
	"testing"
)

func TestResolveOneImportStmt_AlreadyParsed(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "schemas.ts", `export const userSchema = z.object({});`)
	fi := mustParse(t, []byte(`import { userSchema } from './schemas';`))
	fi.Path = filepath.Join(dir, "app.ts")
	ctx := newScanCtx(dir)
	resolved := filepath.Join(dir, "schemas.ts")
	ctx.parsed[resolved] = fi
	resolveOneImportStmt(ctx, fi, firstImportStmt(t, fi), map[string]bool{"userSchema": true})
	if len(ctx.schemas) != 0 {
		t.Fatalf("expected skip when already parsed, got %v", ctx.schemas)
	}
}

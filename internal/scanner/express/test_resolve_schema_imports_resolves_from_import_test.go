//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestResolveSchemaImports_ResolvesFromImport 테스트
package express

import (
	"path/filepath"
	"testing"
)

func TestResolveSchemaImports_ResolvesFromImport(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "schemas.ts", `export const userSchema = z.object({ name: z.string() });`)
	appPath := filepath.Join(dir, "app.ts")
	src := []byte(`import { userSchema } from './schemas';
const r = express.Router();
r.post('/x', validateRequest({ body: userSchema }), h);
`)
	fi := mustParse(t, src)
	fi.Path = appPath
	ctx := newScanCtx(dir)
	ctx.allRouters = map[string]map[string]bool{appPath: {"r": true}}
	ctx.parsed[appPath] = fi

	resolveSchemaImports(ctx)
	if _, ok := ctx.schemas["userSchema"]; !ok {
		t.Fatalf("expected userSchema resolved from import, got %v", ctx.schemas)
	}
}

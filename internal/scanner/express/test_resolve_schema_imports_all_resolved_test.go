//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestResolveSchemaImports_AllResolved 테스트
package express

import (
	"path/filepath"
	"testing"
)

func TestResolveSchemaImports_AllResolved(t *testing.T) {
	dir := t.TempDir()
	appPath := filepath.Join(dir, "app.ts")
	src := []byte(`const r = express.Router();
const userSchema = z.object({ name: z.string() });
r.post('/x', validateRequest({ body: userSchema }), h);
`)
	fi := mustParse(t, src)
	fi.Path = appPath
	ctx := newScanCtx(dir)
	ctx.parsed[appPath] = fi
	ctx.allRouters[appPath] = map[string]bool{"r": true}

	ctx.schemas["userSchema"] = nil

	resolveSchemaImports(ctx)

	if len(ctx.parsed) != 1 {
		t.Fatalf("expected no extra parse, got %d", len(ctx.parsed))
	}
}

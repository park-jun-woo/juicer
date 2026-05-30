//ff:func feature=scan type=test control=sequence topic=express
//ff:what resolveSchemaImports: 미해결 스키마 import 추적 파싱 + 조기반환 분기
package express

import (
	"path/filepath"
	"testing"
)

func TestResolveSchemaImports_NoNeeded(t *testing.T) {
	ctx := newScanCtx(t.TempDir())
	// no parsed files -> no needed schemas -> early return
	resolveSchemaImports(ctx)
	if len(ctx.schemas) != 0 {
		t.Fatalf("expected no schemas, got %v", ctx.schemas)
	}
}

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
	// schema already locally present -> unresolvedSet empty -> early return
	ctx.schemas["userSchema"] = nil

	resolveSchemaImports(ctx)
	// nothing extra parsed
	if len(ctx.parsed) != 1 {
		t.Fatalf("expected no extra parse, got %d", len(ctx.parsed))
	}
}

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

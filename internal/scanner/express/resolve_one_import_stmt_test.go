//ff:func feature=scan type=test control=sequence topic=express
//ff:what resolveOneImportStmt: 미해결 스키마 import 파싱·병합 + 각 조기반환 분기
package express

import (
	"path/filepath"
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
)

func newScanCtx(absRoot string) *scanContext {
	return &scanContext{
		parsed:      map[string]*fileInfo{},
		allRouters:  map[string]map[string]bool{},
		schemas:     map[string]*sitter.Node{},
		schemaSrc:   map[string][]byte{},
		absRoot:     absRoot,
		pathAliases: map[string]string{},
	}
}

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

func TestResolveOneImportStmt_NotNeeded(t *testing.T) {
	dir := t.TempDir()
	fi := mustParse(t, []byte(`import { other } from './schemas';`))
	fi.Path = filepath.Join(dir, "app.ts")
	ctx := newScanCtx(dir)
	resolveOneImportStmt(ctx, fi, firstImportStmt(t, fi), map[string]bool{"userSchema": true})
	if len(ctx.schemas) != 0 {
		t.Fatalf("expected no merge, got %v", ctx.schemas)
	}
}

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

// Remaining uncovered branches (line 17 empty importPath with a needed name,
// line 29 parseFile error after resolveSourceBase confirmed the file) are not
// reachable from valid, parseable TypeScript import statements.

func TestResolveOneImportStmt_AlreadyParsed(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "schemas.ts", `export const userSchema = z.object({});`)
	fi := mustParse(t, []byte(`import { userSchema } from './schemas';`))
	fi.Path = filepath.Join(dir, "app.ts")
	ctx := newScanCtx(dir)
	resolved := filepath.Join(dir, "schemas.ts")
	ctx.parsed[resolved] = fi // mark already parsed
	resolveOneImportStmt(ctx, fi, firstImportStmt(t, fi), map[string]bool{"userSchema": true})
	if len(ctx.schemas) != 0 {
		t.Fatalf("expected skip when already parsed, got %v", ctx.schemas)
	}
}

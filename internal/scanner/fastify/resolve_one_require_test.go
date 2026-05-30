//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what resolveOneRequire 테스트
package fastify

import (
	"os"
	"path/filepath"
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
)

func reqDeclarator(t *testing.T, src string) (*sitter.Node, *fileInfo) {
	t.Helper()
	fi := mustParse(t, []byte(src))
	ds := findAllByType(fi.Root, "variable_declarator")
	if len(ds) == 0 {
		t.Fatalf("no declarator in %q", src)
	}
	return ds[0], fi
}

func TestResolveOneRequire_Resolved(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "mod.ts"), []byte(""), 0o644)
	d, fi := reqDeclarator(t, `const m = require("./mod");`+"\n")
	imports := map[string]string{}
	resolveOneRequire(d, fi.Src, dir, imports)
	if imports["m"] == "" {
		t.Fatalf("expected m resolved, got %v", imports)
	}
}

func TestResolveOneRequire_NotRequire(t *testing.T) {
	d, fi := reqDeclarator(t, `const m = foo("./mod");`+"\n")
	imports := map[string]string{}
	resolveOneRequire(d, fi.Src, t.TempDir(), imports)
	if len(imports) != 0 {
		t.Fatalf("expected no imports for non-require, got %v", imports)
	}
}

func TestResolveOneRequire_NoInitCall(t *testing.T) {
	d, fi := reqDeclarator(t, "const m = 5;\n")
	imports := map[string]string{}
	resolveOneRequire(d, fi.Src, t.TempDir(), imports)
	if len(imports) != 0 {
		t.Fatalf("expected no imports, got %v", imports)
	}
}

func TestResolveOneRequire_NoPath(t *testing.T) {
	d, fi := reqDeclarator(t, "const m = require();\n")
	imports := map[string]string{}
	resolveOneRequire(d, fi.Src, t.TempDir(), imports)
	if len(imports) != 0 {
		t.Fatalf("expected no imports for require with no arg, got %v", imports)
	}
}

func TestResolveOneRequire_IndexResolution(t *testing.T) {
	dir := t.TempDir()
	sub := filepath.Join(dir, "pkg")
	os.MkdirAll(sub, 0o755)
	os.WriteFile(filepath.Join(sub, "index.ts"), []byte(""), 0o644)
	d, fi := reqDeclarator(t, `const p = require("./pkg");`+"\n")
	imports := map[string]string{}
	resolveOneRequire(d, fi.Src, dir, imports)
	if imports["p"] == "" {
		t.Fatalf("expected p resolved to index.ts, got %v", imports)
	}
}

func TestResolveOneRequire_External(t *testing.T) {
	d, fi := reqDeclarator(t, `const fastify = require("fastify");`+"\n")
	imports := map[string]string{}
	resolveOneRequire(d, fi.Src, t.TempDir(), imports)
	if len(imports) != 0 {
		t.Fatalf("external require should not resolve, got %v", imports)
	}
}

//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what resolveImports 테스트
package fastify

import (
	"os"
	"path/filepath"
	"testing"
)

func TestResolveImports(t *testing.T) {
	dir := t.TempDir()
	// a local module to import
	os.WriteFile(filepath.Join(dir, "users.ts"), []byte("export default 1;"), 0o644)

	appPath := filepath.Join(dir, "app.ts")
	src := `
import usersPlugin from "./users";
import Fastify from "fastify";
const cors = require("./users");
`
	fi := mustParse(t, []byte(src))
	fi.Path = appPath

	imports := resolveImports(fi, dir)
	if imports["usersPlugin"] == "" {
		t.Fatalf("expected usersPlugin resolved, got %v", imports)
	}
	// external module "fastify" is not resolved to a local file
	if _, ok := imports["Fastify"]; ok {
		t.Errorf("external import should not resolve: %v", imports)
	}
}

func TestResolveImports_None(t *testing.T) {
	fi := mustParse(t, []byte("const x = 1;\n"))
	fi.Path = "/app/main.ts"
	imports := resolveImports(fi, "/app")
	if len(imports) != 0 {
		t.Fatalf("expected no imports, got %v", imports)
	}
}

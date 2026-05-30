//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestResolveImports 테스트
package fastify

import (
	"os"
	"path/filepath"
	"testing"
)

func TestResolveImports(t *testing.T) {
	dir := t.TempDir()

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

	if _, ok := imports["Fastify"]; ok {
		t.Errorf("external import should not resolve: %v", imports)
	}
}

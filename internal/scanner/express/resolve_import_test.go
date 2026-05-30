//ff:func feature=scan type=test control=sequence topic=express
//ff:what resolveImports: import 및 require 변수→파일 경로 매핑
package express

import (
	"path/filepath"
	"testing"
)

func TestResolveImports_ImportAndRequire(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "users.ts", "x")
	writeFile(t, dir, "admin.ts", "x")

	src := []byte(`import userRouter from './users';
const adminRouter = require('./admin');
`)
	fi := mustParse(t, src)
	fi.Path = filepath.Join(dir, "app.ts")

	imports := resolveImports(fi, dir, nil)
	if got := imports["userRouter"]; got != filepath.Join(dir, "users.ts") {
		t.Errorf("userRouter -> %q", got)
	}
	if got := imports["adminRouter"]; got != filepath.Join(dir, "admin.ts") {
		t.Errorf("adminRouter -> %q", got)
	}
}

func TestResolveImports_Empty(t *testing.T) {
	fi := mustParse(t, []byte(`const x = 1;`))
	fi.Path = "/tmp/app.ts"
	imports := resolveImports(fi, "/tmp", nil)
	if len(imports) != 0 {
		t.Fatalf("expected empty, got %v", imports)
	}
}

//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestResolveOneRequire_Resolves 테스트
package express

import (
	"path/filepath"
	"testing"
)

func TestResolveOneRequire_Resolves(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "users.ts", "x")
	fi := mustParse(t, []byte(`const r = require('./users');`))
	imports := map[string]string{}
	resolveOneRequire(firstLexDecl(t, fi), fi.Src, dir, imports, dir, nil)
	if imports["r"] != filepath.Join(dir, "users.ts") {
		t.Fatalf("got %v", imports)
	}
}

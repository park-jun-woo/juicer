//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestResolveImports_Empty 테스트
package express

import "testing"

func TestResolveImports_Empty(t *testing.T) {
	fi := mustParse(t, []byte(`const x = 1;`))
	fi.Path = "/tmp/app.ts"
	imports := resolveImports(fi, "/tmp", nil)
	if len(imports) != 0 {
		t.Fatalf("expected empty, got %v", imports)
	}
}

//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestResolveOneImport_EmptyPath 테스트
package express

import "testing"

func TestResolveOneImport_EmptyPath(t *testing.T) {

	fi := mustParse(t, []byte(`const x = 1;`))
	imports := map[string]string{}
	resolveOneImport(fi.Root, fi.Src, t.TempDir(), imports, t.TempDir(), nil)
	if len(imports) != 0 {
		t.Fatalf("expected no bindings, got %v", imports)
	}
}

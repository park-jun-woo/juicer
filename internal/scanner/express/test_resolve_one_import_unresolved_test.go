//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestResolveOneImport_Unresolved 테스트
package express

import "testing"

func TestResolveOneImport_Unresolved(t *testing.T) {
	dir := t.TempDir()
	fi := mustParse(t, []byte(`import r from 'external-pkg';`))
	imports := map[string]string{}
	resolveOneImport(firstImportStmt(t, fi), fi.Src, dir, imports, dir, nil)
	if len(imports) != 0 {
		t.Fatalf("expected unresolved, got %v", imports)
	}
}

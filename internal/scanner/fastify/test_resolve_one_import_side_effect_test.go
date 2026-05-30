//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestResolveOneImport_SideEffect 테스트
package fastify

import "testing"

func TestResolveOneImport_SideEffect(t *testing.T) {

	dir := t.TempDir()
	fi := mustParse(t, []byte(`import "./styles.css";`+"\n"))
	stmt := findAllByType(fi.Root, "import_statement")[0]
	imports := map[string]string{}
	resolveOneImport(stmt, fi.Src, dir, imports)
	if len(imports) != 0 {
		t.Fatalf("expected no imports for side-effect, got %v", imports)
	}
}

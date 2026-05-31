//ff:func feature=scan type=test topic=express control=sequence
//ff:what reexportHasBinding export specifier 바인딩명 일치 여부 테스트
package express

import "testing"

func TestReexportHasBinding(t *testing.T) {
	fi := mustParse(t, []byte(`export { foo, bar as baz } from './m';`))
	stmts := findAllByType(fi.Root, "export_statement")
	if len(stmts) == 0 {
		t.Fatal("no export_statement")
	}
	stmt := stmts[0]
	if !reexportHasBinding(stmt, fi.Src, "foo") {
		t.Error("foo binding should match")
	}
	if !reexportHasBinding(stmt, fi.Src, "baz") {
		t.Error("alias baz should match")
	}
	if reexportHasBinding(stmt, fi.Src, "missing") {
		t.Error("missing should be false")
	}
}

//ff:func feature=scan type=test topic=express control=sequence
//ff:what constObjectInDecl 선언 내 이름 일치 object value 노드 반환 테스트
package express

import "testing"

func TestConstObjectInDecl(t *testing.T) {
	fi := mustParse(t, []byte(`const schema = { a: 1 };`))
	decls := findAllByType(fi.Root, "lexical_declaration")
	if len(decls) == 0 {
		t.Fatal("no lexical_declaration")
	}
	if obj := constObjectInDecl(decls[0], fi.Src, "schema"); obj == nil || obj.Type() != "object" {
		t.Errorf("schema object: %v", obj)
	}
	// name mismatch
	if obj := constObjectInDecl(decls[0], fi.Src, "other"); obj != nil {
		t.Error("name mismatch should be nil")
	}
	// non-object value
	fi2 := mustParse(t, []byte(`const n = 5;`))
	d2 := findAllByType(fi2.Root, "lexical_declaration")[0]
	if obj := constObjectInDecl(d2, fi2.Src, "n"); obj != nil {
		t.Error("non-object should be nil")
	}
}

//ff:func feature=scan type=test topic=express control=sequence
//ff:what findConstObject 이름 일치 const object 리터럴 노드 반환 테스트
package express

import "testing"

func TestFindConstObject(t *testing.T) {
	fi := mustParse(t, []byte(`const schema = { a: 1 };
const other = 5;`))
	if obj := findConstObject(fi.Root, fi.Src, "schema"); obj == nil || obj.Type() != "object" {
		t.Errorf("schema object not found: %v", obj)
	}
	// non-object value
	if obj := findConstObject(fi.Root, fi.Src, "other"); obj != nil {
		t.Error("other is not an object")
	}
	// missing name
	if obj := findConstObject(fi.Root, fi.Src, "ghost"); obj != nil {
		t.Error("ghost should be nil")
	}
}

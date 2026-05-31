//ff:func feature=ddl type=test control=sequence
//ff:what sortedEnums 이름 기준 정렬 복사 테스트
package ddl

import "testing"

func TestSortedEnums(t *testing.T) {
	in := []EnumType{{Name: "Role"}, {Name: "Color"}, {Name: "Status"}}
	got := sortedEnums(in)
	if got[0].Name != "Color" || got[1].Name != "Role" || got[2].Name != "Status" {
		t.Errorf("got %+v", got)
	}
	// original untouched
	if in[0].Name != "Role" {
		t.Errorf("original mutated: %+v", in)
	}
}

//ff:func feature=scan type=test topic=nestjs control=iteration dimension=1
//ff:what enumMemberKeyName enum_assignment/property_identifier 키명 추출 테스트
package nestjs

import "testing"

func TestEnumMemberKeyName(t *testing.T) {
	src := []byte(`enum E { Asset = 'assets', Bare }`)
	root, err := parseTypeScript(src)
	if err != nil {
		t.Fatal(err)
	}
	body := findChildByType(findAllByType(root, "enum_declaration")[0], "enum_body")
	if body == nil {
		t.Fatal("no enum_body")
	}
	var names []string
	for i := 0; i < int(body.ChildCount()); i++ {
		if n := enumMemberKeyName(body.Child(i), src); n != "" {
			names = append(names, n)
		}
	}
	if len(names) != 2 || names[0] != "Asset" || names[1] != "Bare" {
		t.Errorf("got %v", names)
	}
}

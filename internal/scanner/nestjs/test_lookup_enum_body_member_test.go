//ff:func feature=scan type=test topic=nestjs control=sequence
//ff:what lookupEnumBodyMember enum_body에서 키 일치 멤버 값 조회 테스트
package nestjs

import "testing"

func TestLookupEnumBodyMember(t *testing.T) {
	src := []byte(`enum E { A = 'a', B = 'b' }`)
	root, err := parseTypeScript(src)
	if err != nil {
		t.Fatal(err)
	}
	body := findChildByType(findAllByType(root, "enum_declaration")[0], "enum_body")
	if v, ok := lookupEnumBodyMember(body, src, "B"); !ok || v != "b" {
		t.Errorf("B: (%q,%v)", v, ok)
	}
	if _, ok := lookupEnumBodyMember(body, src, "Z"); ok {
		t.Error("missing should be false")
	}
}

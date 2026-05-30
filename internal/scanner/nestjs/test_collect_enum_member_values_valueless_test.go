//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestCollectEnumMemberValues_Valueless 테스트
package nestjs

import "testing"

func TestCollectEnumMemberValues_Valueless(t *testing.T) {
	src := []byte(`enum Dir { Up, Down }`)
	root, _ := parseTypeScript(src)
	bodies := findAllByType(root, "enum_body")
	vals := collectEnumMemberValues(bodies[0], src)
	if len(vals) != 2 || vals[0] != "Up" {
		t.Fatalf("got %v", vals)
	}
}

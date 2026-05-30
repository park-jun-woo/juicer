//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestCollectEnumMemberValues_Assignment 테스트
package nestjs

import "testing"

func TestCollectEnumMemberValues_Assignment(t *testing.T) {
	src := []byte(`enum Status { OPEN = 'open', CLOSED = 'closed' }`)
	root, _ := parseTypeScript(src)
	bodies := findAllByType(root, "enum_body")
	if len(bodies) == 0 {
		t.Fatal("no enum_body")
	}
	vals := collectEnumMemberValues(bodies[0], src)
	if len(vals) != 2 || vals[0] != "open" {
		t.Fatalf("got %v", vals)
	}
}

//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestChildrenOfType 테스트
package quarkus

import "testing"

func TestChildrenOfType(t *testing.T) {
	root, _ := parseJava([]byte(`class C { int a; int b; }`))
	body := findAllByType(root, "class_body")[0]
	fields := childrenOfType(body, "field_declaration")
	if len(fields) != 2 {
		t.Fatalf("got %d", len(fields))
	}
}

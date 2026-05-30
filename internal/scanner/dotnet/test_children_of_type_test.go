//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what TestChildrenOfType 테스트
package dotnet

import "testing"

func TestChildrenOfType(t *testing.T) {
	root, _ := parseCS(t, `class C { int a; int b; }`)
	body := findAllByType(root, "declaration_list")[0]
	fields := childrenOfType(body, "field_declaration")
	if len(fields) != 2 {
		t.Fatalf("got %d", len(fields))
	}
}

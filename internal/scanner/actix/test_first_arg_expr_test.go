//ff:func feature=scan type=test topic=actix control=sequence
//ff:what firstArgExpr arguments 노드의 첫 named child 반환 테스트
package actix

import "testing"

func TestFirstArgExpr(t *testing.T) {
	root, err := parseRust([]byte(`fn f() { g(42, 7); }`))
	if err != nil {
		t.Fatal(err)
	}
	args := findAllByType(root, "arguments")
	if len(args) == 0 {
		t.Fatal("no arguments")
	}
	first := firstArgExpr(args[0])
	if first == nil || first.Content([]byte(`fn f() { g(42, 7); }`)) != "42" {
		t.Errorf("first arg: %v", first)
	}
	// empty arguments -> nil
	root2, _ := parseRust([]byte(`fn f() { g(); }`))
	args2 := findAllByType(root2, "arguments")
	if firstArgExpr(args2[0]) != nil {
		t.Error("empty args should be nil")
	}
}

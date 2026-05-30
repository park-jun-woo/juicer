//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestFindFieldReceiver_HasReceiver 테스트
package actix

import "testing"

func TestFindFieldReceiver_HasReceiver(t *testing.T) {
	src := []byte(`fn f() { foo.bar; }`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	fe := firstFieldExpr(root)
	if fe == nil {
		t.Fatal("no field_expression")
	}
	recv := findFieldReceiver(fe)
	if recv == nil {
		t.Fatal("expected non-nil receiver")
	}
	if nodeText(recv, src) != "foo" {
		t.Errorf("receiver = %q, want foo", nodeText(recv, src))
	}
}

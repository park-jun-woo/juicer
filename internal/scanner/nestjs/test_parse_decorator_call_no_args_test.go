//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestParseDecoratorCall_NoArgs 테스트
package nestjs

import "testing"

func TestParseDecoratorCall_NoArgs(t *testing.T) {
	src := []byte(`
@Get()
function f() {}
`)
	root, err := parseTypeScript(src)
	if err != nil {
		t.Fatal(err)
	}
	calls := findAllByType(root, "call_expression")
	if len(calls) == 0 {
		t.Fatal("no call_expression found")
	}
	d := parseDecoratorCall(calls[0], src)
	if d.name != "Get" {
		t.Fatalf("expected Get, got %q", d.name)
	}
}

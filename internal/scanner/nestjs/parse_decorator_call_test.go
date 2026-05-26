//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestParseDecoratorCall_StringArg 테스트
package nestjs

import "testing"

func TestParseDecoratorCall_StringArg(t *testing.T) {
	src := []byte(`
@Get('users')
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
	if d.arg != "users" {
		t.Fatalf("expected users, got %q", d.arg)
	}
}

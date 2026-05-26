//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestParseDecoratorCall_NumberArg 테스트
package nestjs

import "testing"

func TestParseDecoratorCall_NumberArg(t *testing.T) {
	src := []byte(`
@HttpCode(201)
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
	if d.name != "HttpCode" {
		t.Fatalf("expected HttpCode, got %q", d.name)
	}
	if d.arg != "201" {
		t.Fatalf("expected 201, got %q", d.arg)
	}
}

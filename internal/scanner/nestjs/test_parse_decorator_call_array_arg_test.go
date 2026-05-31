//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestParseDecoratorCall_ArrayArg 테스트 (@Get(['/a','/b']) 배열 경로)
package nestjs

import "testing"

func TestParseDecoratorCall_ArrayArg(t *testing.T) {
	src := []byte(`
@Get(['/api/v1/tables', '/api/v2/tables'])
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
	if len(d.args) != 2 {
		t.Fatalf("expected 2 array paths, got %d (%v)", len(d.args), d.args)
	}
	if d.args[0] != "/api/v1/tables" || d.args[1] != "/api/v2/tables" {
		t.Fatalf("unexpected array paths: %v", d.args)
	}
	if d.arg != "/api/v1/tables" {
		t.Fatalf("expected first path as d.arg, got %q", d.arg)
	}
}

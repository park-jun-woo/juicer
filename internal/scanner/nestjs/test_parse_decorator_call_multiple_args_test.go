//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestParseDecoratorCall_MultipleArgs 테스트
package nestjs

import "testing"

func TestParseDecoratorCall_MultipleArgs(t *testing.T) {
	src := []byte(`
@UseGuards(JwtAuthGuard, RolesGuard)
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
	if d.name != "UseGuards" {
		t.Fatalf("expected UseGuards, got %q", d.name)
	}
	if len(d.args) != 2 {
		t.Fatalf("expected 2 args, got %d: %v", len(d.args), d.args)
	}
	if d.args[0] != "JwtAuthGuard" {
		t.Fatalf("expected JwtAuthGuard, got %q", d.args[0])
	}
	if d.args[1] != "RolesGuard" {
		t.Fatalf("expected RolesGuard, got %q", d.args[1])
	}
	if d.arg != "JwtAuthGuard" {
		t.Fatalf("expected arg=JwtAuthGuard, got %q", d.arg)
	}
}

//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestParseDecoratorCall_IdentifierArg 테스트
package nestjs

import "testing"

func TestParseDecoratorCall_IdentifierArg(t *testing.T) {
	src := []byte(`
@UseGuards(JwtAuthGuard)
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
	if d.arg != "JwtAuthGuard" {
		t.Fatalf("expected JwtAuthGuard, got %q", d.arg)
	}
	if len(d.args) != 1 || d.args[0] != "JwtAuthGuard" {
		t.Fatalf("expected args=[JwtAuthGuard], got %v", d.args)
	}
}

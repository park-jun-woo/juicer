//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestParseDecoratorCall_MemberExpressionArg 테스트
package nestjs

import "testing"

func TestParseDecoratorCall_MemberExpressionArg(t *testing.T) {
	src := []byte(`
@Roles(Role.premium)
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
	if d.name != "Roles" {
		t.Fatalf("expected Roles, got %q", d.name)
	}
	if d.arg != "Role.premium" {
		t.Fatalf("expected Role.premium, got %q", d.arg)
	}
}

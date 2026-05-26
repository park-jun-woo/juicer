//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestParseDecoratorCall_ObjectArg 테스트
package nestjs

import "testing"

func TestParseDecoratorCall_ObjectArg(t *testing.T) {
	src := []byte(`
@Controller({ path: 'auth', version: '1' })
export class AuthController {}
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
	if d.name != "Controller" {
		t.Fatalf("expected Controller, got %q", d.name)
	}
	if d.arg != "auth" {
		t.Fatalf("expected auth, got %q", d.arg)
	}
	if d.objectProps == nil {
		t.Fatal("expected objectProps to be non-nil")
	}
	if d.objectProps["version"] != "1" {
		t.Fatalf("expected version 1, got %q", d.objectProps["version"])
	}
	if d.objectProps["path"] != "auth" {
		t.Fatalf("expected path auth, got %q", d.objectProps["path"])
	}
}

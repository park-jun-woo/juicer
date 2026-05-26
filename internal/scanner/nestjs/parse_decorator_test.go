//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what parseDecorator 테스트
package nestjs

import "testing"

func TestParseDecorator_Call(t *testing.T) {
	src := []byte(`
@Controller('users')
export class UsersController {}
`)
	root, err := parseTypeScript(src)
	if err != nil {
		t.Fatal(err)
	}
	decs := findAllByType(root, "decorator")
	if len(decs) == 0 {
		t.Fatal("no decorators found")
	}
	d := parseDecorator(decs[0], src)
	if d.name != "Controller" {
		t.Fatalf("expected Controller, got %q", d.name)
	}
	if d.arg != "users" {
		t.Fatalf("expected users, got %q", d.arg)
	}
}

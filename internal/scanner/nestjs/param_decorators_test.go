//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what paramDecorators 테스트
package nestjs

import "testing"

func TestParamDecorators_Body(t *testing.T) {
	src := []byte(`
class C {
  create(@Body() dto: CreateUserDto) {}
}
`)
	root, err := parseTypeScript(src)
	if err != nil {
		t.Fatal(err)
	}
	params := findAllByType(root, "required_parameter")
	if len(params) == 0 {
		t.Fatal("no params")
	}
	decs := paramDecorators(params[0], src)
	if len(decs) == 0 {
		t.Fatal("expected decorators")
	}
	if decs[0].name != "Body" {
		t.Fatalf("expected Body, got %q", decs[0].name)
	}
}

//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what extractOneParam 테스트
package nestjs

import "testing"

func TestExtractOneParam_Body(t *testing.T) {
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
	result := &methodParams{}
	extractOneParam(params[0], src, result)
	if result.bodyType != "CreateUserDto" {
		t.Fatalf("expected CreateUserDto, got %q", result.bodyType)
	}
}

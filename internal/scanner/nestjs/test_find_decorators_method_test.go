//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestFindDecorators_Method 테스트
package nestjs

import "testing"

func TestFindDecorators_Method(t *testing.T) {
	src := []byte(`
class C {
  @Get()
  findAll() {}
}
`)
	root, err := parseTypeScript(src)
	if err != nil {
		t.Fatal(err)
	}
	methods := findAllByType(root, "method_definition")
	if len(methods) == 0 {
		t.Fatal("no methods")
	}
	decs := findDecorators(methods[0], src)
	if len(decs) == 0 {
		t.Fatal("expected decorators")
	}
}

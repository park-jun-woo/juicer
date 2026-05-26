//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what collectPrecedingSiblingDecorators 테스트
package nestjs

import "testing"

func TestCollectPrecedingSiblingDecorators_Found(t *testing.T) {
	src := []byte(`
class C {
  @Get()
  @HttpCode(200)
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
	m := methods[0]
	parent := m.Parent()
	decs := collectPrecedingSiblingDecorators(parent, m, src)
	if len(decs) < 1 {
		t.Fatal("expected decorators")
	}
}

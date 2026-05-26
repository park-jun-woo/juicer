//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestFindFactoryInnerClass_Found 테스트
package nestjs

import "testing"

func TestFindFactoryInnerClass_Found(t *testing.T) {
	src := []byte(`
export function BaseController(CreateDto, UpdateDto) {
  class GenericsController {
    @Get()
    findAll() {}
  }
  return GenericsController;
}
`)
	root, err := parseTypeScript(src)
	if err != nil {
		t.Fatal(err)
	}
	cls := findFactoryInnerClass(root, src, "BaseController")
	if cls == nil {
		t.Fatal("expected inner class, got nil")
	}
	nameNode := findChildByType(cls, "type_identifier")
	if nameNode == nil {
		t.Fatal("inner class has no name")
	}
	if got := nodeText(nameNode, src); got != "GenericsController" {
		t.Fatalf("expected GenericsController, got %q", got)
	}
}

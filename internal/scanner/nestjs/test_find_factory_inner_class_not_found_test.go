//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestFindFactoryInnerClass_NotFound 테스트
package nestjs

import "testing"

func TestFindFactoryInnerClass_NotFound(t *testing.T) {
	src := []byte(`
export function OtherFactory() {
  return {};
}
`)
	root, err := parseTypeScript(src)
	if err != nil {
		t.Fatal(err)
	}
	cls := findFactoryInnerClass(root, src, "BaseController")
	if cls != nil {
		t.Fatal("expected nil, got a node")
	}
}

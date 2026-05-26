//ff:func feature=scan type=test control=iteration dimension=1 topic=nestjs
//ff:what TestExtractOneMethod_NoHTTPDecorator 테스트
package nestjs

import "testing"

func TestExtractOneMethod_NoHTTPDecorator(t *testing.T) {
	src := []byte(`
class Svc {
  doSomething() {}
}
`)
	root, err := parseTypeScript(src)
	if err != nil {
		t.Fatal(err)
	}
	methods := findAllByType(root, "method_definition")
	for _, m := range methods {
		_, ok := extractOneMethod(m, src, "test.ts")
		if ok {
			t.Fatal("expected not ok for non-HTTP method")
		}
	}
}

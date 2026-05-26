//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestExtractReturnType_NoAnnotation 테스트
package nestjs

import "testing"

func TestExtractReturnType_NoAnnotation(t *testing.T) {
	src := []byte(`
class C {
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
	rt := extractReturnType(methods[0], src)
	if rt != "" {
		t.Fatalf("expected empty, got %q", rt)
	}
}

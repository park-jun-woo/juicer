//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what extractTypeAnnotation 테스트
package nestjs

import "testing"

func TestExtractTypeAnnotation_Basic(t *testing.T) {
	src := []byte(`function f(x: string) {}`)
	root, err := parseTypeScript(src)
	if err != nil {
		t.Fatal(err)
	}
	anns := findAllByType(root, "type_annotation")
	if len(anns) == 0 {
		t.Fatal("no type_annotation")
	}
	typ := extractTypeAnnotation(anns[0], src)
	if typ != "string" {
		t.Fatalf("expected string, got %q", typ)
	}
}

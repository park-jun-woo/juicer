//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestExtractClassAttribute_NotFound 테스트
package django

import "testing"

func TestExtractClassAttribute_NotFound(t *testing.T) {
	src := `
class C:
    other = X
`
	root, err := parsePython([]byte(src))
	if err != nil {
		t.Fatal(err)
	}
	body := firstBlock(root)
	if got := extractClassAttribute(body, "serializer_class", []byte(src)); got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}

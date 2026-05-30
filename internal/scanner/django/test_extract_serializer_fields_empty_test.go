//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestExtractSerializerFields_Empty 테스트
package django

import "testing"

func TestExtractSerializerFields_Empty(t *testing.T) {
	src := `
class C:
    helper = 42
`
	root, err := parsePython([]byte(src))
	if err != nil {
		t.Fatal(err)
	}
	body := firstBlock(root)
	if f := extractSerializerFields(body, []byte(src)); len(f) != 0 {
		t.Fatalf("expected no serializer fields, got %+v", f)
	}
}

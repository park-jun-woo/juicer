//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestExtractFieldsList_NoFields 테스트
package django

import "testing"

func TestExtractFieldsList_NoFields(t *testing.T) {

	src := `
class Meta:
    model = User
    exclude = ('secret',)
    foo()
`
	root, err := parsePython([]byte(src))
	if err != nil {
		t.Fatal(err)
	}
	body := firstBlock(root)
	if f := extractFieldsList(body, []byte(src)); f != nil {
		t.Fatalf("expected nil, got %v", f)
	}
}

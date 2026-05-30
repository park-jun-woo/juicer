//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestExtractFieldsList_FieldsNotList 테스트
package django

import "testing"

func TestExtractFieldsList_FieldsNotList(t *testing.T) {
	src := `
class Meta:
    fields = '__all__'
`
	root, err := parsePython([]byte(src))
	if err != nil {
		t.Fatal(err)
	}
	body := firstBlock(root)
	if f := extractFieldsList(body, []byte(src)); f != nil {
		t.Fatalf("expected nil for non-list fields, got %v", f)
	}
}

//ff:func feature=scan type=test control=sequence topic=django
//ff:what extractFieldsList — Meta body의 fields=[...] 추출 분기를 검증
package django

import "testing"

func TestExtractFieldsList(t *testing.T) {
	src := `
class Meta:
    model = User
    fields = ['id', 'name']
`
	root, err := parsePython([]byte(src))
	if err != nil {
		t.Fatal(err)
	}
	body := firstBlock(root)
	if body == nil {
		t.Fatal("no Meta body")
	}
	fields := extractFieldsList(body, []byte(src))
	if len(fields) != 2 || fields[0] != "id" || fields[1] != "name" {
		t.Fatalf("expected [id name], got %v", fields)
	}
}

func TestExtractFieldsList_NoFields(t *testing.T) {
	// Has assignments but none named "fields"; also a non-list and a call stmt.
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

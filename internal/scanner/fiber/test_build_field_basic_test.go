//ff:func feature=scan type=test control=sequence
//ff:what TestBuildField_Basic 테스트
package fiber

import "testing"

func TestBuildField_Basic(t *testing.T) {
	src := `package m
type Book struct {
	Title  string ` + "`json:\"title\"`" + `
	Hidden string ` + "`json:\"-\"`" + `
}
`
	st, _ := structFields(t, src, "Book")

	f0 := buildField(st.Field(0), st.Tag(0), map[string]bool{})
	if f0 == nil || f0.Name != "Title" || f0.JSON != "title" {
		t.Fatalf("Title field = %+v", f0)
	}
	if f0.Type != "string" {
		t.Errorf("Title type = %q", f0.Type)
	}

	f1 := buildField(st.Field(1), st.Tag(1), map[string]bool{})
	if f1 != nil {
		t.Fatalf("json:\"-\" field should be nil, got %+v", f1)
	}
}

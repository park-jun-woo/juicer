//ff:func feature=scan type=test control=sequence
//ff:what TestBuildField_NoTag 테스트
package fiber

import "testing"

func TestBuildField_NoTag(t *testing.T) {
	src := `package m
type T struct {
	Plain int
}
`
	st, _ := structFields(t, src, "T")
	f := buildField(st.Field(0), st.Tag(0), map[string]bool{})
	if f == nil || f.Name != "Plain" {
		t.Fatalf("Plain field = %+v", f)
	}
}

//ff:func feature=scan type=test control=sequence
//ff:what TestResolveNestedFields_ArrayOfStruct 테스트
package fiber

import "testing"

func TestResolveNestedFields_ArrayOfStruct(t *testing.T) {
	src := `package m
type E struct { N int ` + "`json:\"n\"`" + ` }
type A struct { Items [3]E }
`
	st, _ := structFields(t, src, "A")
	got := resolveNestedFields(st.Field(0).Type(), map[string]bool{})
	if len(got) != 1 || got[0].Name != "N" {
		t.Fatalf("array of struct nested = %v", got)
	}
}

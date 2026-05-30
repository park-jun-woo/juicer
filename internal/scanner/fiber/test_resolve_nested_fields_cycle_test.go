//ff:func feature=scan type=test control=sequence
//ff:what TestResolveNestedFields_Cycle 테스트
package fiber

import "testing"

func TestResolveNestedFields_Cycle(t *testing.T) {
	src := `package m
type S struct { X int }
`
	st, _ := structFields(t, src, "S")
	typ := st.Field(0).Type()
	visited := map[string]bool{typ.String(): true}
	if got := resolveNestedFields(typ, visited); got != nil {
		t.Fatalf("visited should yield nil, got %v", got)
	}
}

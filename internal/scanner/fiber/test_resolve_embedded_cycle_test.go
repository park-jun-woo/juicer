//ff:func feature=scan type=test control=sequence
//ff:what TestResolveEmbedded_Cycle 테스트
package fiber

import "testing"

func TestResolveEmbedded_Cycle(t *testing.T) {
	src := `package m
type T struct { X int }
`
	st, _ := structFields(t, src, "T")
	typ := st.Field(0).Type()

	visited := map[string]bool{typ.String(): true}
	if got := resolveEmbedded(typ, visited); got != nil {
		t.Fatalf("expected nil for visited, got %v", got)
	}
}

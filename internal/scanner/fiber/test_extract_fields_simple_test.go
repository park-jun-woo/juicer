//ff:func feature=scan type=test control=sequence
//ff:what TestExtractFields_Simple 테스트
package fiber

import "testing"

func TestExtractFields_Simple(t *testing.T) {
	src := `package m
type T struct {
	A string
	B int
}
`
	st, _ := structFields(t, src, "T")
	fields := extractFields(st, map[string]bool{})
	if len(fields) != 2 {
		t.Fatalf("expected 2 fields, got %d", len(fields))
	}
}

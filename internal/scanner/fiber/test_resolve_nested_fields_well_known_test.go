//ff:func feature=scan type=test control=sequence
//ff:what TestResolveNestedFields_WellKnown 테스트
package fiber

import "testing"

func TestResolveNestedFields_WellKnown(t *testing.T) {
	src := `package m
import "time"
type T struct { When time.Time }
`
	st, _ := structFields(t, src, "T")
	if got := resolveNestedFields(st.Field(0).Type(), map[string]bool{}); got != nil {
		t.Fatalf("time.Time should not expand, got %v", got)
	}
}

//ff:func feature=scan type=test control=sequence
//ff:what TestWellKnownType_Custom 테스트
package fiber

import "testing"

func TestWellKnownType_Custom(t *testing.T) {
	src := `package m
type Custom struct{}
var C Custom
`
	named := namedTypeOf(t, src, "C")
	if _, ok := wellKnownType(named); ok {
		t.Fatal("custom type should not be well-known")
	}
}

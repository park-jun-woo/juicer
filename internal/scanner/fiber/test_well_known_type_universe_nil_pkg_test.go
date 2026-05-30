//ff:func feature=scan type=test control=sequence
//ff:what TestWellKnownType_UniverseNilPkg 테스트
package fiber

import "testing"

func TestWellKnownType_UniverseNilPkg(t *testing.T) {
	src := `package m
var E error
`
	named := namedTypeOf(t, src, "E")

	if _, ok := wellKnownType(named); ok {
		t.Fatal("error (universe) should not be well-known")
	}
}

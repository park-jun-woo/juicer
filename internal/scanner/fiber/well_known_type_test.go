//ff:func feature=scan type=test control=selection
//ff:what wellKnownType — well-known 타입 판정 테스트
package fiber

import (
	"go/types"
	"testing"
)

func namedTypeOf(t *testing.T, src, varName string) *types.Named {
	t.Helper()
	typ := typeOfVar(t, src, varName)
	if ptr, ok := typ.(*types.Pointer); ok {
		typ = ptr.Elem()
	}
	named, ok := typ.(*types.Named)
	if !ok {
		t.Fatalf("%s is not a named type: %T", varName, typ)
	}
	return named
}

func TestWellKnownType_TimeTime(t *testing.T) {
	src := `package m
import "time"
var T time.Time
`
	named := namedTypeOf(t, src, "T")
	if name, ok := wellKnownType(named); !ok || name != "time.Time" {
		t.Fatalf("time.Time: %q %v", name, ok)
	}
}

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

func TestWellKnownType_UniverseNilPkg(t *testing.T) {
	src := `package m
var E error
`
	named := namedTypeOf(t, src, "E")
	// error's package is nil (universe scope) -> false
	if _, ok := wellKnownType(named); ok {
		t.Fatal("error (universe) should not be well-known")
	}
}

//ff:func feature=scan type=extract control=iteration dimension=1
//ff:what TestIsIntKind 테스트
package scanner

import (
	"go/types"
	"testing"
)

func TestIsIntKind(t *testing.T) {
	intKinds := []types.BasicKind{
		types.Int, types.Int8, types.Int16, types.Int32, types.Int64,
		types.Uint, types.Uint8, types.Uint16, types.Uint32, types.Uint64,
		types.Uintptr,
	}
	for _, k := range intKinds {
		if !isIntKind(k) {
			t.Errorf("expected isIntKind(%v) = true", k)
		}
	}

	nonIntKinds := []types.BasicKind{
		types.Bool, types.String, types.Float32, types.Float64,
		types.Complex64, types.Complex128,
	}
	for _, k := range nonIntKinds {
		if isIntKind(k) {
			t.Errorf("expected isIntKind(%v) = false", k)
		}
	}
}

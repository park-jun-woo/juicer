//ff:func feature=scan type=test control=selection
//ff:what isIntKind — 정수 계열 BasicKind 판정 테스트
package fiber

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
			t.Errorf("expected true for kind %v", k)
		}
	}
	for _, k := range []types.BasicKind{types.String, types.Bool, types.Float64, types.Complex128} {
		if isIntKind(k) {
			t.Errorf("expected false for kind %v", k)
		}
	}
}

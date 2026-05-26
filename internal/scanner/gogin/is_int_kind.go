//ff:func feature=scan type=extract control=selection
//ff:what types.BasicKind가 정수 계열인지 검사한다
package gogin

import (
	"go/types"
)

func isIntKind(k types.BasicKind) bool {
	switch k {
	case types.Int, types.Int8, types.Int16, types.Int32, types.Int64,
		types.Uint, types.Uint8, types.Uint16, types.Uint32, types.Uint64,
		types.Uintptr:
		return true
	}
	return false
}


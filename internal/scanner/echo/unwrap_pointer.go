//ff:func feature=scan type=extract control=sequence
//ff:what 포인터 타입을 벗겨서 내부 타입을 반환한다
package echo

import (
	"go/types"
)

func unwrapPointer(t types.Type) types.Type {
	if ptr, ok := t.(*types.Pointer); ok {
		return ptr.Elem()
	}
	return t
}

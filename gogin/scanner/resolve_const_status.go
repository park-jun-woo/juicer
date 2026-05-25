//ff:func feature=scan type=extract control=sequence
//ff:what types.Object가 상수이면 그 값을 문자열로 반환한다
package scanner

import "go/types"

// resolveConstStatus returns the string representation of a constant object, or empty string.
func resolveConstStatus(obj types.Object) string {
	c, ok := obj.(*types.Const)
	if !ok {
		return ""
	}
	return constToString(c.Val())
}

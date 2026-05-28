//ff:func feature=scan type=extract control=sequence
//ff:what types.Object가 상수이면 그 값을 문자열로 반환한다
package fiber

import (
	"go/types"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

// resolveConstStatus returns the string representation of a constant object, or empty string.
func resolveConstStatus(obj types.Object) string {
	c, ok := obj.(*types.Const)
	if !ok {
		return ""
	}
	return scanner.ConstToString(c.Val())
}

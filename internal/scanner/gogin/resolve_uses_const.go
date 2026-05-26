//ff:func feature=scan type=extract control=sequence
//ff:what info.Uses에서 식별자가 상수이면 그 값을 문자열로 반환한다
package gogin

import (
	"go/ast"
	"go/types"
)

// resolveUsesConst looks up an identifier in Uses and returns its const value as string.
func resolveUsesConst(info *types.Info, ident *ast.Ident) string {
	obj, ok := info.Uses[ident]
	if !ok {
		return ""
	}
	return resolveConstStatus(obj)
}

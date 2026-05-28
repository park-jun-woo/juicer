//ff:func feature=scan type=extract control=iteration dimension=1
//ff:what 호출 인자 중 라우터 파라미터명과 일치하는 Ident의 인덱스를 반환한다
package fiber

import (
	"go/ast"
)

func routerArgIndex(call *ast.CallExpr, paramName string) int {
	for i, arg := range call.Args {
		if id, ok := arg.(*ast.Ident); ok && id.Name == paramName {
			return i
		}
	}
	return -1
}

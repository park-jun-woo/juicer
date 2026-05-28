//ff:func feature=scan type=extract control=sequence
//ff:what types.Type이 echo 라우터 타입(*echo.Echo, *echo.Group)인지 검사한다
package echo

import (
	"go/types"
	"strings"
)

func isEchoRouterTypeInfo(t types.Type) bool {
	var named *types.Named
	if ptr, ok := t.(*types.Pointer); ok {
		named, _ = ptr.Elem().(*types.Named)
	} else {
		named, _ = t.(*types.Named)
	}
	if named == nil {
		return false
	}
	obj := named.Obj()
	if !echoRouterTypes[obj.Name()] {
		return false
	}
	pkg := obj.Pkg()
	if pkg == nil {
		return false
	}
	return strings.HasSuffix(pkg.Path(), "labstack/echo/v4")
}

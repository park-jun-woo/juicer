//ff:func feature=scan type=extract control=sequence
//ff:what types.Type이 echo.Context인지 검사한다
package echo

import (
	"go/types"
	"strings"
)

func isEchoContextTypeInfo(t types.Type) bool {
	// Echo Context는 interface이므로 포인터가 아닐 수 있음
	// 하지만 구현체 포인터로 전달될 수 있으므로 포인터도 확인
	if ptr, ok := t.(*types.Pointer); ok {
		t = ptr.Elem()
	}

	named, ok := t.(*types.Named)
	if !ok {
		return false
	}
	obj := named.Obj()
	if obj.Name() != "Context" {
		return false
	}
	pkg := obj.Pkg()
	if pkg == nil {
		return false
	}
	return strings.HasSuffix(pkg.Path(), "labstack/echo/v4")
}

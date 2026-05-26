//ff:func feature=scan type=extract control=sequence
//ff:what types.Type이 *gin.RouterGroup 또는 *gin.Engine인지 검사한다
package gogin

import (
	"go/types"
	"strings"
)

func isGinRouterTypeInfo(t types.Type) bool {
	ptr, ok := t.(*types.Pointer)
	if !ok {
		return false
	}
	named, ok := ptr.Elem().(*types.Named)
	if !ok {
		return false
	}
	obj := named.Obj()
	if !ginRouterTypes[obj.Name()] {
		return false
	}
	pkg := obj.Pkg()
	if pkg == nil {
		return false
	}
	return strings.HasSuffix(pkg.Path(), "gin-gonic/gin")
}

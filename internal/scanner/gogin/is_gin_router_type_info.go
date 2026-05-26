//ff:func feature=scan type=extract control=sequence
//ff:what types.Type이 gin 라우터 타입(*gin.Engine, *gin.RouterGroup, gin.IRouter 등)인지 검사한다
package gogin

import (
	"go/types"
	"strings"
)

func isGinRouterTypeInfo(t types.Type) bool {
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
	if !ginRouterTypes[obj.Name()] {
		return false
	}
	pkg := obj.Pkg()
	if pkg == nil {
		return false
	}
	return strings.HasSuffix(pkg.Path(), "gin-gonic/gin")
}

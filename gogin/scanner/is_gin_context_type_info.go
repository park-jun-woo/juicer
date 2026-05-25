//ff:func feature=scan type=extract control=sequence
//ff:what types.Type이 *gin.Context인지 검사한다
package scanner

import (
	"go/types"
	"strings"
)

func isGinContextTypeInfo(t types.Type) bool {
	ptr, ok := t.(*types.Pointer)
	if !ok {
		return false
	}
	named, ok := ptr.Elem().(*types.Named)
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
	return strings.HasSuffix(pkg.Path(), "gin-gonic/gin")
}


//ff:func feature=scan type=extract control=sequence
//ff:what types.Type이 *fiber.Ctx인지 검사한다
package fiber

import (
	"go/types"
	"strings"
)

func isFiberContextTypeInfo(t types.Type) bool {
	ptr, ok := t.(*types.Pointer)
	if !ok {
		return false
	}
	named, ok := ptr.Elem().(*types.Named)
	if !ok {
		return false
	}
	obj := named.Obj()
	if obj.Name() != "Ctx" {
		return false
	}
	pkg := obj.Pkg()
	if pkg == nil {
		return false
	}
	return strings.HasPrefix(pkg.Path(), "github.com/gofiber/fiber")
}

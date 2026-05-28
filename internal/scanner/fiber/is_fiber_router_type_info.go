//ff:func feature=scan type=extract control=sequence
//ff:what types.Type이 fiber 라우터 타입(*fiber.App, *fiber.Group, fiber.Router)인지 검사한다
package fiber

import (
	"go/types"
	"strings"
)

func isFiberRouterTypeInfo(t types.Type) bool {
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
	if !fiberRouterTypes[obj.Name()] {
		return false
	}
	pkg := obj.Pkg()
	if pkg == nil {
		return false
	}
	return strings.HasPrefix(pkg.Path(), "github.com/gofiber/fiber")
}

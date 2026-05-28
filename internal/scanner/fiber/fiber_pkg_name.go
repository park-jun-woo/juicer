//ff:func feature=scan type=extract control=iteration dimension=2
//ff:what 파일의 fiber import alias를 결정한다
package fiber

import (
	"go/ast"
	"strings"
)

func fiberPkgName(file *ast.File) string {
	for _, imp := range file.Imports {
		p := strings.Trim(imp.Path.Value, `"`)
		if p == fiberPkgPath || strings.HasPrefix(p, "github.com/gofiber/fiber") {
			if imp.Name != nil {
				return imp.Name.Name
			}
			// Default import name for "github.com/gofiber/fiber/v2" is "fiber"
			return "fiber"
		}
	}
	return ""
}

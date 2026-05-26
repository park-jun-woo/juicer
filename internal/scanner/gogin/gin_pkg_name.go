//ff:func feature=scan type=extract control=iteration dimension=2
//ff:what ginPkgName 함수
package gogin

import (
	"go/ast"
	"strings"
)

func ginPkgName(file *ast.File) string {
	for _, imp := range file.Imports {
		p := strings.Trim(imp.Path.Value, `"`)
		if p == "github.com/gin-gonic/gin" {
			if imp.Name != nil {
				return imp.Name.Name
			}
			return "gin"
		}
	}
	return ""
}

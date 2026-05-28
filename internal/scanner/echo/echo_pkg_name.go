//ff:func feature=scan type=extract control=iteration dimension=2
//ff:what 파일의 echo import alias를 결정한다
package echo

import (
	"go/ast"
	"strings"
)

func echoPkgName(file *ast.File) string {
	for _, imp := range file.Imports {
		p := strings.Trim(imp.Path.Value, `"`)
		if p == echoPkgPath {
			if imp.Name != nil {
				return imp.Name.Name
			}
			return "echo"
		}
	}
	return ""
}

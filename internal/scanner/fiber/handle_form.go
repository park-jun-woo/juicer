//ff:func feature=scan type=extract control=iteration dimension=1
//ff:what c.FormValue("name") 호출에서 폼 필드를 추출한다
package fiber

import (
	"go/ast"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func handleForm(ep *scanner.Endpoint, call *ast.CallExpr) {
	if len(call.Args) < 1 {
		return
	}
	name := stringLitValue(call.Args[0])
	if name == "" {
		return
	}
	scanner.EnsureRequest(ep)

	for _, f := range ep.Request.FormFields {
		if f.Name == name {
			return
		}
	}
	ep.Request.FormFields = append(ep.Request.FormFields, scanner.Param{Name: name, Type: "string"})
}

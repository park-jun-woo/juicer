//ff:func feature=scan type=extract control=iteration dimension=1
//ff:what c.PostForm("name") 호출에서 폼 필드를 추출한다
package scanner

import (
	"go/ast"
)

func handleForm(ep *Endpoint, call *ast.CallExpr) {
	if len(call.Args) < 1 {
		return
	}
	name := stringLitValue(call.Args[0])
	if name == "" {
		return
	}
	ensureRequest(ep)

	for _, f := range ep.Request.FormFields {
		if f.Name == name {
			return
		}
	}
	ep.Request.FormFields = append(ep.Request.FormFields, Param{Name: name, Type: "string"})
}


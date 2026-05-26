//ff:func feature=scan type=extract control=iteration dimension=1
//ff:what c.FormFile("name") 호출에서 파일 업로드 필드를 추출한다
package gogin

import (
	"go/ast"
	"github.com/park-jun-woo/juicer/internal/scanner"
)

func handleFile(ep *scanner.Endpoint, call *ast.CallExpr) {
	if len(call.Args) < 1 {
		return
	}
	name := stringLitValue(call.Args[0])
	if name == "" {
		return
	}
	scanner.EnsureRequest(ep)

	for _, f := range ep.Request.Files {
		if f.Name == name {
			return
		}
	}
	ep.Request.Files = append(ep.Request.Files, scanner.Param{Name: name, Type: "file"})
}


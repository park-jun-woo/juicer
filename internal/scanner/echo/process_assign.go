//ff:func feature=scan type=extract control=iteration dimension=3
//ff:what 대입문에서 echo.New() 또는 Group 호출을 감지하여 라우터를 등록한다
package echo

import (
	"go/ast"
	"github.com/park-jun-woo/codistill/internal/scanner"
)

func processAssign(stmt *ast.AssignStmt, echoAlias string, routers map[string]*routerInfo) {
	for i, rhs := range stmt.Rhs {
		if i >= len(stmt.Lhs) {
			break
		}
		call, ok := rhs.(*ast.CallExpr)
		if !ok {
			continue
		}
		sel, ok := call.Fun.(*ast.SelectorExpr)
		if !ok {
			continue
		}
		lhs, ok := stmt.Lhs[i].(*ast.Ident)
		if !ok {
			continue
		}

		if isEchoInit(sel, echoAlias) {
			routers[lhs.Name] = &routerInfo{}
			continue
		}

		if sel.Sel.Name == "Group" {
			recv := identName(sel.X)
			parent, ok := routers[recv]
			if !ok {
				continue
			}
			prefix := parent.prefix
			if len(call.Args) > 0 {
				if s, ok := extractPathString(call.Args[0]); ok {
					prefix = scanner.JoinPath(prefix, s)
				}
			}
			mw := append([]string{}, parent.middleware...)
			routers[lhs.Name] = &routerInfo{prefix: prefix, middleware: mw}
		}
	}
}

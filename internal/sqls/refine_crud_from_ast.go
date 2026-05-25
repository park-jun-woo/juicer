//ff:func feature=sql type=parse control=sequence
//ff:what ExecContext 인라인 문자열에서 CRUD 판별
package sqls

import (
	"go/ast"
	"go/token"
	"strings"
)

// refineCRUDFromAST checks inline string literals passed directly to ExecContext.
//
func refineCRUDFromAST(body *ast.BlockStmt) string {
	if body == nil {
		return "EXEC"
	}

	var crud string
	ast.Inspect(body, func(n ast.Node) bool {
		call, ok := n.(*ast.CallExpr)
		if !ok {
			return true
		}
		sel, ok := call.Fun.(*ast.SelectorExpr)
		if !ok {
			return true
		}
		if sel.Sel.Name != "ExecContext" {
			return true
		}
		// Check string arguments
		for _, arg := range call.Args {
			lit, ok := arg.(*ast.BasicLit)
			if !ok || lit.Kind != token.STRING {
				continue
			}
			val := lit.Value
			upper := strings.ToUpper(val)
			if strings.Contains(upper, "INSERT") {
				crud = "INSERT"
			} else if strings.Contains(upper, "UPDATE") && strings.Contains(upper, "SET") {
				crud = "UPDATE"
			} else if strings.Contains(upper, "DELETE") {
				crud = "DELETE"
			}
		}
		return true
	})
	if crud == "" {
		return "EXEC"
	}
	return crud
}


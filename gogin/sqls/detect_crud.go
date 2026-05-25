//ff:func feature=sql type=parse control=sequence
//ff:what 메서드 body에서 DB 호출 패턴으로 CRUD 종류 판별
package sqls

import (
	"go/ast"
)

// detectCRUD inspects a method body for r.db.QueryContext, r.db.QueryRowContext,
// or r.db.ExecContext calls. Returns "SELECT", "EXEC", or "" if none found.
//
func detectCRUD(body *ast.BlockStmt) string {
	if body == nil {
		return ""
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
		switch sel.Sel.Name {
		case "QueryContext":
			crud = "SELECT"
		case "QueryRowContext":
			if crud == "" {
				crud = "SELECT"
			}
		case "ExecContext":
			if crud == "" {
				crud = "EXEC"
			}
		}
		return true
	})
	return crud
}


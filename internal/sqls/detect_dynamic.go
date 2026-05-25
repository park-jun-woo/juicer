//ff:func feature=sql type=parse control=sequence
//ff:what 메서드 body에서 동적 SQL 패턴(+=, fmt.Sprintf) 감지
package sqls

import (
	"go/ast"
	"go/token"
)

// detectDynamic checks if a method body contains += assignments or fmt.Sprintf calls.
//
func detectDynamic(body *ast.BlockStmt) bool {
	if body == nil {
		return false
	}

	dynamic := false
	ast.Inspect(body, func(n ast.Node) bool {
		if dynamic {
			return false
		}
		switch node := n.(type) {
		case *ast.AssignStmt:
			if node.Tok == token.ADD_ASSIGN {
				dynamic = true
				return false
			}
		case *ast.CallExpr:
			sel, ok := node.Fun.(*ast.SelectorExpr)
			if !ok {
				return true
			}
			pkg, ok := sel.X.(*ast.Ident)
			if !ok {
				return true
			}
			if pkg.Name == "fmt" && sel.Sel.Name == "Sprintf" {
				dynamic = true
				return false
			}
		}
		return true
	})
	return dynamic
}


//ff:func feature=sql type=parse control=selection
//ff:what 메서드 리시버에서 타입명 추출 (*T, T 모두 처리)
package sqls

import (
	"go/ast"
)

// receiverTypeName extracts the type name from a method receiver,
// handling both *T and T forms.
//
func receiverTypeName(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.StarExpr:
		if id, ok := t.X.(*ast.Ident); ok {
			return id.Name
		}
	case *ast.Ident:
		return t.Name
	}
	return ""
}


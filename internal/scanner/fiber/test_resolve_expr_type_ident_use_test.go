//ff:func feature=scan type=test control=sequence
//ff:what TestResolveExprType_IdentUse 테스트
package fiber

import (
	"go/ast"
	"testing"
)

func TestResolveExprType_IdentUse(t *testing.T) {
	src := `package m
type Resp struct { OK bool ` + "`json:\"ok\"`" + ` }
func h() {
	var r Resp
	use(r)
}
func use(interface{}) {}
`
	file, info := typedExprs(t, src)
	var arg ast.Expr
	ast.Inspect(file, func(n ast.Node) bool {
		if c, ok := n.(*ast.CallExpr); ok {
			if id, ok := c.Fun.(*ast.Ident); ok && id.Name == "use" {
				arg = c.Args[0]
			}
		}
		return true
	})
	tn, fields := resolveExprType(arg, info)
	if tn != "Resp" || len(fields) != 1 {
		t.Fatalf("ident use: %q %v", tn, fields)
	}
}

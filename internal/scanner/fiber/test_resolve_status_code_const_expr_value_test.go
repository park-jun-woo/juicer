//ff:func feature=scan type=test control=sequence
//ff:what TestResolveStatusCode_ConstExprValue 테스트
package fiber

import (
	"go/ast"
	"testing"
)

func TestResolveStatusCode_ConstExprValue(t *testing.T) {

	src := `package m
func use(int) {}
func h() { use(200 + 1) }
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
	if got := resolveStatusCode(arg, info); got != "201" {
		t.Fatalf("200+1 -> %q, want 201", got)
	}
}

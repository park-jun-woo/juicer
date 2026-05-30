//ff:func feature=scan type=test control=sequence
//ff:what TestResolveStatusCode_ConstIdent 테스트
package fiber

import (
	"go/ast"
	"testing"
)

func TestResolveStatusCode_ConstIdent(t *testing.T) {
	src := `package m
const Created = 201
func use(int) {}
func h() { use(Created) }
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
		t.Fatalf("Created -> %q, want 201", got)
	}
}

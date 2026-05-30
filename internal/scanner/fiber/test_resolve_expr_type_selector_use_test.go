//ff:func feature=scan type=test control=sequence
//ff:what TestResolveExprType_SelectorUse 테스트
package fiber

import (
	"go/ast"
	"testing"
)

func TestResolveExprType_SelectorUse(t *testing.T) {

	src := `package m
type Inner struct { Z int ` + "`json:\"z\"`" + ` }
type Outer struct { In Inner }
func use(interface{}) {}
func h() {
	var o Outer
	use(o.In)
}
`
	file, info := typedExprs(t, src)
	var arg *ast.SelectorExpr
	ast.Inspect(file, func(n ast.Node) bool {
		if c, ok := n.(*ast.CallExpr); ok {
			if id, ok := c.Fun.(*ast.Ident); ok && id.Name == "use" {
				if s, ok := c.Args[0].(*ast.SelectorExpr); ok {
					arg = s
				}
			}
		}
		return true
	})
	if arg == nil {
		t.Fatal("selector arg not found")
	}
	tn, _ := resolveExprType(arg, info)
	if tn != "Inner" {
		t.Fatalf("selector use: %q, want Inner", tn)
	}
}

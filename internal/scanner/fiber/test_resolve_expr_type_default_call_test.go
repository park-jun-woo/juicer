//ff:func feature=scan type=test control=sequence
//ff:what TestResolveExprType_DefaultCall 테스트
package fiber

import (
	"go/ast"
	"testing"
)

func TestResolveExprType_DefaultCall(t *testing.T) {
	src := `package m
type R struct { V int ` + "`json:\"v\"`" + ` }
func make2() R { return R{} }
func use(interface{}) {}
func h() { use(make2()) }
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
	tn, _ := resolveExprType(arg, info)
	if tn != "R" {
		t.Fatalf("default call: %q, want R", tn)
	}
}

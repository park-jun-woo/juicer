//ff:func feature=scan type=test control=sequence
//ff:what TestResolveUsesConst_Found 테스트
package fiber

import (
	"go/ast"
	"testing"
)

func TestResolveUsesConst_Found(t *testing.T) {
	src := `package m
const Created = 201
func use(int) {}
func h() { use(Created) }
`
	file, info := typedExprs(t, src)
	var ident *ast.Ident
	ast.Inspect(file, func(n ast.Node) bool {
		if c, ok := n.(*ast.CallExpr); ok {
			if id, ok := c.Fun.(*ast.Ident); ok && id.Name == "use" {
				ident = c.Args[0].(*ast.Ident)
			}
		}
		return true
	})
	if got := resolveUsesConst(info, ident); got != "201" {
		t.Fatalf("const use = %q, want 201", got)
	}
}

//ff:func feature=scan type=test control=sequence
//ff:what TestResolveResponseType_NoFields 테스트
package fiber

import (
	"go/ast"
	"testing"
)

func TestResolveResponseType_NoFields(t *testing.T) {

	src := `package m
func use(interface{}) {}
func h() { var n int; use(n) }
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
	_, fields, conf := resolveResponseType(arg, info)
	if len(fields) != 0 || conf != "" {
		t.Fatalf("no fields: %v %q", fields, conf)
	}
}

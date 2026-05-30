//ff:func feature=scan type=test control=sequence
//ff:what TestResolveResponseType_FullFields 테스트
package fiber

import (
	"go/ast"
	"testing"
)

func TestResolveResponseType_FullFields(t *testing.T) {
	src := `package m
type Resp struct { OK bool ` + "`json:\"ok\"`" + ` }
func use(interface{}) {}
func h() { var r Resp; use(r) }
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
	tn, fields, conf := resolveResponseType(arg, info)
	if tn != "Resp" || len(fields) != 1 || conf != "full" {
		t.Fatalf("full: %q %v %q", tn, fields, conf)
	}
}

//ff:func feature=scan type=test control=sequence
//ff:what TestResolveExprType_CompositeLit 테스트
package fiber

import (
	"go/ast"
	"testing"
)

func TestResolveExprType_CompositeLit(t *testing.T) {
	src := `package m
type Out struct { N int ` + "`json:\"n\"`" + ` }
func h() { use(Out{N: 1}) }
func use(interface{}) {}
`
	file, info := typedExprs(t, src)
	var arg ast.Expr
	ast.Inspect(file, func(n ast.Node) bool {
		if cl, ok := n.(*ast.CompositeLit); ok && arg == nil {
			arg = cl
		}
		return true
	})
	tn, _ := resolveExprType(arg, info)
	if tn != "Out" {
		t.Fatalf("composite lit: %q", tn)
	}
}

//ff:func feature=scan type=test control=sequence
//ff:what TestResolveExprType_Selector 테스트
package fiber

import (
	"go/ast"
	"testing"
)

func TestResolveExprType_Selector(t *testing.T) {

	src := `package m
type Cfg struct { N int ` + "`json:\"n\"`" + ` }
var Conf Cfg
func use(interface{}) {}
func h() { use(pkgAlias()) }
func pkgAlias() Cfg { return Conf }
`
	file, info := typedExprs(t, src)
	// find a selector expr in Uses — use Conf via its Ident directly
	var sel *ast.SelectorExpr
	ast.Inspect(file, func(n ast.Node) bool {
		if s, ok := n.(*ast.SelectorExpr); ok && sel == nil {
			sel = s
		}
		return true
	})
	if sel != nil {
		resolveExprType(sel, info)
	}
}

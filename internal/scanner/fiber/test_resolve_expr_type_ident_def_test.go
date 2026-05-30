//ff:func feature=scan type=test control=sequence
//ff:what TestResolveExprType_IdentDef 테스트
package fiber

import (
	"go/ast"
	"testing"
)

func TestResolveExprType_IdentDef(t *testing.T) {

	src := `package m
type D struct { A int ` + "`json:\"a\"`" + ` }
var Decl D
`
	file, info := typedExprs(t, src)
	var defIdent *ast.Ident
	ast.Inspect(file, func(n ast.Node) bool {
		if id, ok := n.(*ast.Ident); ok && id.Name == "Decl" {
			defIdent = id
		}
		return true
	})
	if defIdent == nil {
		t.Fatal("Decl ident not found")
	}
	tn, _ := resolveExprType(defIdent, info)
	if tn != "D" {
		t.Fatalf("ident def: %q, want D", tn)
	}
}

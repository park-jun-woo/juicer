//ff:func feature=scan type=test control=sequence topic=echo
//ff:what TestResolveUsesConst 테스트
package echo

import (
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"testing"
)

func TestResolveUsesConst(t *testing.T) {
	src := `package m
const StatusOK = 200
var _ = StatusOK
`
	fset := token.NewFileSet()
	file, _ := parser.ParseFile(fset, "m.go", src, 0)
	conf := types.Config{Importer: importer.Default()}
	info := &types.Info{
		Defs: map[*ast.Ident]types.Object{},
		Uses: map[*ast.Ident]types.Object{},
	}
	if _, err := conf.Check("m", fset, []*ast.File{file}, info); err != nil {
		t.Fatal(err)
	}
	// find the use of StatusOK
	var useIdent *ast.Ident
	ast.Inspect(file, func(n ast.Node) bool {
		if id, ok := n.(*ast.Ident); ok && id.Name == "StatusOK" {
			if _, isUse := info.Uses[id]; isUse {
				useIdent = id
			}
		}
		return true
	})
	if useIdent == nil {
		t.Fatal("no use found")
	}
	if got := resolveUsesConst(info, useIdent); got != "200" {
		t.Fatalf("got %q", got)
	}
}

//ff:func feature=scan type=test control=sequence
//ff:what resolveConstStatus — 상수 상태 코드 해석 테스트
package fiber

import (
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"testing"
)

func TestResolveConstStatus_Const(t *testing.T) {
	src := `package m
const StatusOK = 200
var V = 5
`
	fset := token.NewFileSet()
	file, _ := parser.ParseFile(fset, "m.go", src, 0)
	conf := types.Config{Importer: importer.Default()}
	info := &types.Info{Defs: map[*ast.Ident]types.Object{}}
	pkg, err := conf.Check("m", fset, []*ast.File{file}, info)
	if err != nil {
		t.Fatal(err)
	}
	constObj := pkg.Scope().Lookup("StatusOK")
	if got := resolveConstStatus(constObj); got != "200" {
		t.Fatalf("const status = %q, want 200", got)
	}

	// a non-const object -> ""
	varObj := pkg.Scope().Lookup("V")
	if got := resolveConstStatus(varObj); got != "" {
		t.Fatalf("non-const should be empty, got %q", got)
	}
}

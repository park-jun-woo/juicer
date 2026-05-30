//ff:func feature=scan type=test control=sequence topic=echo
//ff:what TestLookupFunc_Round5 테스트
package echo

import (
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"testing"
)

func TestLookupFunc_Round5(t *testing.T) {
	src := `package m
func Target() {}
`
	fset := token.NewFileSet()
	file, _ := parser.ParseFile(fset, "m.go", src, parser.SkipObjectResolution)
	fn := file.Decls[0].(*ast.FuncDecl)
	info := &types.Info{}
	idx := &funcIndex{
		byPos: map[token.Pos]*ast.FuncDecl{fn.Pos(): fn},
		info:  map[token.Pos]*types.Info{fn.Pos(): info},
	}
	gotFn, gotInfo := lookupFunc(fn.Pos(), idx)
	if gotFn != fn || gotInfo != info {
		t.Fatalf("lookup mismatch")
	}

	if f, _ := lookupFunc(token.NoPos, idx); f != nil {
		t.Fatal("unknown pos should be nil")
	}
}

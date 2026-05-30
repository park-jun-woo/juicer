//ff:func feature=scan type=test control=sequence topic=echo
//ff:what TestParamFieldAtIndex_Round5 테스트
package echo

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestParamFieldAtIndex_Round5(t *testing.T) {
	src := `package m
func F(a int, b, c string) {}
`
	fset := token.NewFileSet()
	file, _ := parser.ParseFile(fset, "m.go", src, 0)
	fn := file.Decls[0].(*ast.FuncDecl)
	_, name0 := paramFieldAtIndex(fn.Type.Params, 0)
	if name0 != "a" {
		t.Fatalf("idx0: %q", name0)
	}
	_, name2 := paramFieldAtIndex(fn.Type.Params, 2)
	if name2 != "c" {
		t.Fatalf("idx2: %q", name2)
	}
	if f, _ := paramFieldAtIndex(fn.Type.Params, 9); f != nil {
		t.Fatal("out-of-range should be nil")
	}
}

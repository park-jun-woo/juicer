//ff:func feature=scan type=test topic=fiber control=sequence
//ff:what astStructFields struct → exported 필드 슬라이스 변환 테스트
package fiber

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestAstStructFields(t *testing.T) {
	f, err := parser.ParseFile(token.NewFileSet(), "x.go", "package p\ntype T struct {\n A string\n B int\n c bool\n}", 0)
	if err != nil {
		t.Fatal(err)
	}
	var st *ast.StructType
	ast.Inspect(f, func(n ast.Node) bool {
		if s, ok := n.(*ast.StructType); ok && st == nil {
			st = s
		}
		return true
	})
	if st == nil {
		t.Fatal("no struct")
	}
	got := astStructFields(st)
	if len(got) != 2 || got[0].Name != "A" || got[1].Name != "B" {
		t.Errorf("got %+v", got)
	}
}

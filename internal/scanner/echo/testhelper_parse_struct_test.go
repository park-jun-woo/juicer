//ff:func feature=scan type=test topic=echo control=sequence
//ff:what echo 테스트 헬퍼 — Go 소스 파싱 후 첫 StructType 노드 반환
package echo

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func parseEchoStruct(t *testing.T, src string) *ast.StructType {
	t.Helper()
	f, err := parser.ParseFile(token.NewFileSet(), "x.go", src, 0)
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
	return st
}

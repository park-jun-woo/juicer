//ff:func feature=scan type=test topic=echo control=sequence
//ff:what indexFileDecls 파일 선언 순회로 함수/struct 인덱싱 테스트
package echo

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestIndexFileDecls(t *testing.T) {
	f, err := parser.ParseFile(token.NewFileSet(), "x.go", `package p
type Dto struct { A int }
func Handler() {}`, 0)
	if err != nil {
		t.Fatal(err)
	}
	idx := &funcIndex{
		byName:     map[string]*ast.FuncDecl{},
		byPos:      map[token.Pos]*ast.FuncDecl{},
		astStructs: map[string]*ast.StructType{},
	}
	indexFileDecls(f, nil, idx)
	if _, ok := idx.byName["Handler"]; !ok {
		t.Errorf("func not indexed: %v", idx.byName)
	}
	if _, ok := idx.astStructs["Dto"]; !ok {
		t.Errorf("struct not indexed: %v", idx.astStructs)
	}
}

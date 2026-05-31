//ff:func feature=scan type=test topic=echo control=iteration dimension=1
//ff:what indexFuncDecl FuncDecl을 byName 인덱스에 등록(비함수/무body 무시) 테스트
package echo

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestIndexFuncDecl(t *testing.T) {
	f, err := parser.ParseFile(token.NewFileSet(), "x.go", `package p
func Handler() {}
func NoBody()
var X = 1`, parser.AllErrors)
	if err != nil {
		// NoBody() without body is a parse error in a normal file; build minimal AST instead.
		f, err = parser.ParseFile(token.NewFileSet(), "x.go", `package p
func Handler() {}
var X = 1`, 0)
		if err != nil {
			t.Fatal(err)
		}
	}
	idx := &funcIndex{
		byName:     map[string]*ast.FuncDecl{},
		byPos:      map[token.Pos]*ast.FuncDecl{},
	}
	for _, decl := range f.Decls {
		indexFuncDecl(decl, nil, idx)
	}
	if _, ok := idx.byName["Handler"]; !ok {
		t.Errorf("Handler not indexed: %v", idx.byName)
	}
	// var decl (non-func) is ignored — no panic, no entry
	if _, ok := idx.byName["X"]; ok {
		t.Error("var X must not be indexed")
	}
}

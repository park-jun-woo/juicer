//ff:func feature=scan type=test topic=fiber control=iteration dimension=1
//ff:what indexFuncDecl FuncDecl byName 등록(비함수/무body 무시) 직접 테스트
package fiber

import (
	"go/parser"
	"go/token"
	"testing"
)

func TestIndexFuncDecl(t *testing.T) {
	f, err := parser.ParseFile(token.NewFileSet(), "x.go", `package p
func Handler() {}
var X = 1`, 0)
	if err != nil {
		t.Fatal(err)
	}
	idx := newFiberIndex()
	for _, decl := range f.Decls {
		indexFuncDecl(decl, nil, idx)
	}
	if _, ok := idx.byName["Handler"]; !ok {
		t.Errorf("Handler not indexed: %v", idx.byName)
	}
	if _, ok := idx.byName["X"]; ok {
		t.Error("var must not index")
	}
}

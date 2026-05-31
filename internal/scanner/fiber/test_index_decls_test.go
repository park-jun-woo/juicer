//ff:func feature=scan type=test topic=fiber control=sequence
//ff:what indexFileDecls/indexFuncDecl/registerStructSpec 파일 선언 인덱싱 테스트
package fiber

import (
	"go/parser"
	"go/token"
	"testing"
)

func TestIndexFileDecls(t *testing.T) {
	f, err := parser.ParseFile(token.NewFileSet(), "x.go", `package p
type Dto struct { A int }
type Alias = int
func Handler() {}
var X = 1`, 0)
	if err != nil {
		t.Fatal(err)
	}
	idx := newFiberIndex()
	indexFileDecls(f, nil, idx)
	if _, ok := idx.byName["Handler"]; !ok {
		t.Errorf("func not indexed: %v", idx.byName)
	}
	if _, ok := idx.astStructs["Dto"]; !ok {
		t.Errorf("struct not indexed: %v", idx.astStructs)
	}
	if _, ok := idx.astStructs["Alias"]; ok {
		t.Error("non-struct alias must not register")
	}
	if _, ok := idx.byName["X"]; ok {
		t.Error("var must not register")
	}
}

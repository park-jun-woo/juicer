//ff:func feature=scan type=test topic=echo control=iteration dimension=1
//ff:what registerStructSpec TypeSpec struct를 인덱스에 등록(비struct/중복 무시) 테스트
package echo

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestRegisterStructSpec(t *testing.T) {
	f, err := parser.ParseFile(token.NewFileSet(), "x.go", `package p
type T struct { A int }
type Alias = int`, 0)
	if err != nil {
		t.Fatal(err)
	}
	idx := &funcIndex{astStructs: map[string]*ast.StructType{}}
	for _, decl := range f.Decls {
		gd, ok := decl.(*ast.GenDecl)
		if !ok {
			continue
		}
		for _, spec := range gd.Specs {
			registerStructSpec(spec, idx)
		}
	}
	if _, ok := idx.astStructs["T"]; !ok {
		t.Errorf("T struct not registered: %v", idx.astStructs)
	}
	if _, ok := idx.astStructs["Alias"]; ok {
		t.Error("non-struct Alias must not register")
	}
}

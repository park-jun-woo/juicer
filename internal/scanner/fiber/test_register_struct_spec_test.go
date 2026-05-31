//ff:func feature=scan type=test topic=fiber control=iteration dimension=1
//ff:what registerStructSpec TypeSpec struct 등록(비struct 무시) 직접 테스트
package fiber

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
	idx := newFiberIndex()
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
		t.Errorf("T not registered: %v", idx.astStructs)
	}
	if _, ok := idx.astStructs["Alias"]; ok {
		t.Error("alias must not register")
	}
}

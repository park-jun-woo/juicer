//ff:func feature=scan type=test control=sequence
//ff:what buildFuncIndex — 함수 인덱스 구축 테스트
package fiber

import (
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"testing"

	"golang.org/x/tools/go/packages"
)

func TestBuildFuncIndex_Empty(t *testing.T) {
	idx := buildFuncIndex(nil)
	if idx == nil || idx.byPos == nil || idx.byName == nil || idx.astStructs == nil {
		t.Fatal("expected initialized index maps")
	}
	if len(idx.byPos) != 0 {
		t.Errorf("expected empty byPos, got %d", len(idx.byPos))
	}
}

func TestBuildFuncIndex_WithPackage(t *testing.T) {
	src := `package m
func Handler() {}
type Book struct{ Title string }
`
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "m.go", src, 0)
	if err != nil {
		t.Fatal(err)
	}
	pkg := &packages.Package{
		Syntax:    []*ast.File{file},
		TypesInfo: &types.Info{},
	}
	idx := buildFuncIndex([]*packages.Package{pkg})
	if idx.byName["Handler"] == nil {
		t.Errorf("expected Handler indexed by name, got %v", idx.byName)
	}
	if _, ok := idx.astStructs["Book"]; !ok {
		t.Errorf("expected Book struct indexed, got %v", idx.astStructs)
	}
}

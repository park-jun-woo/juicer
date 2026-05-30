//ff:func feature=scan type=test control=sequence
//ff:what buildField — 구조체 필드 → Field 변환 테스트
package fiber

import (
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"testing"
)

// structFields type-checks src and returns the named struct's fields as *types.Var.
func structFields(t *testing.T, src, typeName string) (*types.Struct, *types.Info) {
	t.Helper()
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "m.go", src, 0)
	if err != nil {
		t.Fatal(err)
	}
	conf := types.Config{Importer: importer.Default()}
	info := &types.Info{
		Defs:  map[*ast.Ident]types.Object{},
		Types: map[ast.Expr]types.TypeAndValue{},
	}
	pkg, err := conf.Check("m", fset, []*ast.File{file}, info)
	if err != nil {
		t.Fatal(err)
	}
	obj := pkg.Scope().Lookup(typeName)
	if obj == nil {
		t.Fatalf("type %s not found", typeName)
	}
	st, ok := obj.Type().Underlying().(*types.Struct)
	if !ok {
		t.Fatalf("%s is not a struct", typeName)
	}
	return st, info
}

func TestBuildField_Basic(t *testing.T) {
	src := `package m
type Book struct {
	Title  string ` + "`json:\"title\"`" + `
	Hidden string ` + "`json:\"-\"`" + `
}
`
	st, _ := structFields(t, src, "Book")

	// field 0: Title
	f0 := buildField(st.Field(0), st.Tag(0), map[string]bool{})
	if f0 == nil || f0.Name != "Title" || f0.JSON != "title" {
		t.Fatalf("Title field = %+v", f0)
	}
	if f0.Type != "string" {
		t.Errorf("Title type = %q", f0.Type)
	}

	// field 1: Hidden with json:"-" -> nil
	f1 := buildField(st.Field(1), st.Tag(1), map[string]bool{})
	if f1 != nil {
		t.Fatalf("json:\"-\" field should be nil, got %+v", f1)
	}
}

func TestBuildField_NoTag(t *testing.T) {
	src := `package m
type T struct {
	Plain int
}
`
	st, _ := structFields(t, src, "T")
	f := buildField(st.Field(0), st.Tag(0), map[string]bool{})
	if f == nil || f.Name != "Plain" {
		t.Fatalf("Plain field = %+v", f)
	}
}

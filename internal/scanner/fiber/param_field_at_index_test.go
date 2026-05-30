//ff:func feature=scan type=test control=iteration dimension=1
//ff:what paramFieldAtIndex — 파라미터 인덱스 조회 테스트
package fiber

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func paramsOf(t *testing.T, src string) *ast.FieldList {
	t.Helper()
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "m.go", src, 0)
	if err != nil {
		t.Fatal(err)
	}
	for _, d := range file.Decls {
		if fn, ok := d.(*ast.FuncDecl); ok {
			return fn.Type.Params
		}
	}
	t.Fatal("no func")
	return nil
}

func TestParamFieldAtIndex_Named(t *testing.T) {
	params := paramsOf(t, "package m\nfunc f(a int, b string) {}\n")
	_, name0 := paramFieldAtIndex(params, 0)
	if name0 != "a" {
		t.Errorf("idx 0 = %q", name0)
	}
	_, name1 := paramFieldAtIndex(params, 1)
	if name1 != "b" {
		t.Errorf("idx 1 = %q", name1)
	}
}

func TestParamFieldAtIndex_GroupedNames(t *testing.T) {
	// "a, b int" is a single field with two names
	params := paramsOf(t, "package m\nfunc f(a, b int) {}\n")
	_, name1 := paramFieldAtIndex(params, 1)
	if name1 != "b" {
		t.Errorf("grouped idx 1 = %q", name1)
	}
}

func TestParamFieldAtIndex_Unnamed(t *testing.T) {
	// unnamed param -> synthetic "_"
	params := paramsOf(t, "package m\nfunc f(int) {}\n")
	_, name := paramFieldAtIndex(params, 0)
	if name != "_" {
		t.Errorf("unnamed idx 0 = %q, want _", name)
	}
}

func TestParamFieldAtIndex_OutOfRange(t *testing.T) {
	params := paramsOf(t, "package m\nfunc f(a int) {}\n")
	field, name := paramFieldAtIndex(params, 5)
	if field != nil || name != "" {
		t.Errorf("out of range should be nil,'', got %v %q", field, name)
	}
}

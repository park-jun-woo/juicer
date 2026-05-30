//ff:func feature=scan type=test control=sequence
//ff:what fiberRouterParamAtIndex — 라우터 파라미터 인덱스 테스트
package fiber

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func funcDeclFrom(t *testing.T, src string) *ast.FuncDecl {
	t.Helper()
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "m.go", src, 0)
	if err != nil {
		t.Fatal(err)
	}
	for _, d := range file.Decls {
		if fn, ok := d.(*ast.FuncDecl); ok {
			return fn
		}
	}
	t.Fatal("no func decl")
	return nil
}

func TestFiberRouterParamAtIndex_NilParams(t *testing.T) {
	fn := &ast.FuncDecl{Type: &ast.FuncType{}}
	if got := fiberRouterParamAtIndex(fn, newEmptyInfo(), 0); got != "" {
		t.Fatalf("nil params: got %q", got)
	}
}

func TestFiberRouterParamAtIndex_IndexOutOfRange(t *testing.T) {
	fn := funcDeclFrom(t, "package m\nfunc Setup(app int) {}\n")
	if got := fiberRouterParamAtIndex(fn, newEmptyInfo(), 5); got != "" {
		t.Fatalf("out of range: got %q", got)
	}
}

func TestFiberRouterParamAtIndex_NilInfo(t *testing.T) {
	fn := funcDeclFrom(t, "package m\nfunc Setup(app int) {}\n")
	if got := fiberRouterParamAtIndex(fn, nil, 0); got != "" {
		t.Fatalf("nil info: got %q", got)
	}
}

func TestFiberRouterParamAtIndex_NonRouterType(t *testing.T) {
	// info present but TypeOf returns nil for synthetic node -> "" (no match)
	fn := funcDeclFrom(t, "package m\nfunc Setup(app int) {}\n")
	if got := fiberRouterParamAtIndex(fn, newEmptyInfo(), 0); got != "" {
		t.Fatalf("non-router: got %q", got)
	}
}

func TestFiberRouterParamAtIndex_SecondParam(t *testing.T) {
	// second param index, still non-router type -> ""
	fn := funcDeclFrom(t, "package m\nfunc Setup(a int, b string) {}\n")
	if got := fiberRouterParamAtIndex(fn, newEmptyInfo(), 1); got != "" {
		t.Fatalf("second param: got %q", got)
	}
}

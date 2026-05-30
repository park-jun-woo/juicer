//ff:func feature=scan type=test control=iteration dimension=1
//ff:what registerParams — 함수 파라미터 라우터 등록 테스트
package fiber

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func funcDecl(t *testing.T, src string) *ast.FuncDecl {
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
	t.Fatal("no func")
	return nil
}

func TestRegisterParams_RouterParam(t *testing.T) {
	fn := funcDecl(t, "package m\nfunc Setup(app *fiber.App, other int) {}\n")
	routers := map[string]*routerInfo{}
	registerParams(fn, "fiber", routers)
	if _, ok := routers["app"]; !ok {
		t.Fatalf("expected app registered, got %v", routers)
	}
	if _, ok := routers["other"]; ok {
		t.Fatalf("non-router param should not be registered: %v", routers)
	}
}

func TestRegisterParams_NilParams(t *testing.T) {
	fn := &ast.FuncDecl{Type: &ast.FuncType{}}
	routers := map[string]*routerInfo{}
	registerParams(fn, "fiber", routers)
	if len(routers) != 0 {
		t.Fatalf("expected none, got %v", routers)
	}
}

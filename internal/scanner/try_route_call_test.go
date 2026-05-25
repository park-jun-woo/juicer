package scanner

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestTryRouteCall_NonSelector(t *testing.T) {
	call := &ast.CallExpr{Fun: &ast.Ident{Name: "f"}}
	_, ok := tryRouteCall(call, nil, "", nil)
	if ok {
		t.Fatal("expected false")
	}
}

func TestTryRouteCall_NonMethod(t *testing.T) {
	call := &ast.CallExpr{
		Fun: &ast.SelectorExpr{X: &ast.Ident{Name: "r"}, Sel: &ast.Ident{Name: "Render"}},
	}
	_, ok := tryRouteCall(call, map[string]*routerInfo{}, "", nil)
	if ok {
		t.Fatal("expected false")
	}
}

func TestTryRouteCall_MissingRouter(t *testing.T) {
	call := &ast.CallExpr{
		Fun:  &ast.SelectorExpr{X: &ast.Ident{Name: "r"}, Sel: &ast.Ident{Name: "GET"}},
		Args: []ast.Expr{&ast.BasicLit{Kind: token.STRING, Value: `"/test"`}, &ast.Ident{Name: "h"}},
	}
	_, ok := tryRouteCall(call, map[string]*routerInfo{}, "", nil)
	if ok {
		t.Fatal("expected false for unknown router")
	}
}

func TestTryRouteCall_Valid(t *testing.T) {
	call := &ast.CallExpr{
		Fun:  &ast.SelectorExpr{X: &ast.Ident{Name: "r"}, Sel: &ast.Ident{Name: "GET"}},
		Args: []ast.Expr{&ast.BasicLit{Kind: token.STRING, Value: `"/test"`}, &ast.Ident{Name: "handler"}},
	}
	routers := map[string]*routerInfo{"r": {}}
	fset := token.NewFileSet()
	ep, ok := tryRouteCall(call, routers, "main.go", fset)
	if !ok {
		t.Fatal("expected true")
	}
	if ep.Method != "GET" || ep.Path != "/test" {
		t.Fatalf("unexpected ep: %+v", ep)
	}
}

func TestTryRouteCall_InsufficientArgs(t *testing.T) {
	call := &ast.CallExpr{
		Fun:  &ast.SelectorExpr{X: &ast.Ident{Name: "r"}, Sel: &ast.Ident{Name: "GET"}},
		Args: []ast.Expr{&ast.BasicLit{Kind: token.STRING, Value: `"/test"`}},
	}
	routers := map[string]*routerInfo{"r": {}}
	_, ok := tryRouteCall(call, routers, "", nil)
	if ok {
		t.Fatal("expected false with too few args")
	}
}

//ff:func feature=scan type=test control=iteration dimension=1
//ff:what resolveCallerArgs — caller 인자 매핑 해석 테스트
package fiber

import (
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"testing"
)

func TestResolveCallerArgs_NilParamsOrInfo(t *testing.T) {
	fn := &ast.FuncDecl{Type: &ast.FuncType{}}
	status, tn, f, _ := resolveCallerArgs(fn, &ast.CallExpr{}, newEmptyInfo(), newEmptyInfo())
	if status != "" || tn != "" || f != nil {
		t.Fatalf("nil params: %q %q %v", status, tn, f)
	}
	// nil calleeInfo
	fn2 := &ast.FuncDecl{Type: &ast.FuncType{Params: &ast.FieldList{}}}
	if s, _, _, _ := resolveCallerArgs(fn2, &ast.CallExpr{}, newEmptyInfo(), nil); s != "" {
		t.Fatalf("nil calleeInfo should return empty")
	}
}

func TestResolveCallerArgs_StatusFromInt(t *testing.T) {
	src := `package m
func respond(c interface{}, status int) {}
func h() {
	respond(nil, 201)
}
`
	fset := token.NewFileSet()
	file, _ := parser.ParseFile(fset, "m.go", src, 0)
	conf := types.Config{Importer: importer.Default()}
	info := &types.Info{
		Types: map[ast.Expr]types.TypeAndValue{},
		Defs:  map[*ast.Ident]types.Object{},
		Uses:  map[*ast.Ident]types.Object{},
	}
	if _, err := conf.Check("m", fset, []*ast.File{file}, info); err != nil {
		t.Fatal(err)
	}

	var fnDecl *ast.FuncDecl
	var call *ast.CallExpr
	for _, d := range file.Decls {
		if fn, ok := d.(*ast.FuncDecl); ok && fn.Name.Name == "respond" {
			fnDecl = fn
		}
	}
	ast.Inspect(file, func(n ast.Node) bool {
		if c, ok := n.(*ast.CallExpr); ok {
			if id, ok := c.Fun.(*ast.Ident); ok && id.Name == "respond" {
				call = c
			}
		}
		return true
	})
	if fnDecl == nil || call == nil {
		t.Fatal("respond decl/call not found")
	}
	status, _, _, _ := resolveCallerArgs(fnDecl, call, info, info)
	if status != "201" {
		t.Fatalf("expected status 201, got %q", status)
	}
}

func TestResolveCallerArgs_UnnamedParamAndResponse(t *testing.T) {
	// callee has an unnamed interface{} param and named int param;
	// caller passes a struct (resp) and a status -> both resolved.
	src := `package m
type Out struct {
	V int ` + "`json:\"v\"`" + `
}
func write(interface{}, status int) {}
func h() {
	var o Out
	write(o, 200)
}
`
	fset := token.NewFileSet()
	file, _ := parser.ParseFile(fset, "m.go", src, 0)
	conf := types.Config{Importer: importer.Default()}
	info := &types.Info{
		Types: map[ast.Expr]types.TypeAndValue{},
		Defs:  map[*ast.Ident]types.Object{},
		Uses:  map[*ast.Ident]types.Object{},
	}
	if _, err := conf.Check("m", fset, []*ast.File{file}, info); err != nil {
		t.Fatal(err)
	}
	var fnDecl *ast.FuncDecl
	var call *ast.CallExpr
	for _, d := range file.Decls {
		if fn, ok := d.(*ast.FuncDecl); ok && fn.Name.Name == "write" {
			fnDecl = fn
		}
	}
	ast.Inspect(file, func(n ast.Node) bool {
		if c, ok := n.(*ast.CallExpr); ok {
			if id, ok := c.Fun.(*ast.Ident); ok && id.Name == "write" {
				call = c
			}
		}
		return true
	})
	status, tn, _, _ := resolveCallerArgs(fnDecl, call, info, info)
	if status != "200" {
		t.Errorf("status = %q, want 200", status)
	}
	if tn != "Out" {
		t.Errorf("typeName = %q, want Out", tn)
	}
}

func TestResolveCallerArgs_NamedParamsDefs(t *testing.T) {
	// named params resolved via calleeInfo.Defs
	src := `package m
func send(code int) {}
func h() { send(404) }
`
	fset := token.NewFileSet()
	file, _ := parser.ParseFile(fset, "m.go", src, 0)
	conf := types.Config{Importer: importer.Default()}
	info := &types.Info{
		Types: map[ast.Expr]types.TypeAndValue{},
		Defs:  map[*ast.Ident]types.Object{},
		Uses:  map[*ast.Ident]types.Object{},
	}
	if _, err := conf.Check("m", fset, []*ast.File{file}, info); err != nil {
		t.Fatal(err)
	}
	var fnDecl *ast.FuncDecl
	var call *ast.CallExpr
	for _, d := range file.Decls {
		if fn, ok := d.(*ast.FuncDecl); ok && fn.Name.Name == "send" {
			fnDecl = fn
		}
	}
	ast.Inspect(file, func(n ast.Node) bool {
		if c, ok := n.(*ast.CallExpr); ok {
			if id, ok := c.Fun.(*ast.Ident); ok && id.Name == "send" {
				call = c
			}
		}
		return true
	})
	status, _, _, _ := resolveCallerArgs(fnDecl, call, info, info)
	if status != "404" {
		t.Fatalf("expected 404, got %q", status)
	}
}

func TestResolveCallerArgs_FewerArgsThanParams(t *testing.T) {
	src := `package m
func need(a int, b int) {}
func h() { need(200) }
`
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "m.go", src, parser.AllErrors)
	if err != nil {
		// the call is intentionally arity-mismatched; ignore type errors
		_ = err
	}
	conf := types.Config{Importer: importer.Default(), Error: func(error) {}}
	info := &types.Info{
		Types: map[ast.Expr]types.TypeAndValue{},
		Defs:  map[*ast.Ident]types.Object{},
		Uses:  map[*ast.Ident]types.Object{},
	}
	conf.Check("m", fset, []*ast.File{file}, info)

	var fnDecl *ast.FuncDecl
	var call *ast.CallExpr
	for _, d := range file.Decls {
		if fn, ok := d.(*ast.FuncDecl); ok && fn.Name.Name == "need" {
			fnDecl = fn
		}
	}
	ast.Inspect(file, func(n ast.Node) bool {
		if c, ok := n.(*ast.CallExpr); ok {
			if id, ok := c.Fun.(*ast.Ident); ok && id.Name == "need" {
				call = c
			}
		}
		return true
	})
	// 2 params but 1 arg -> n truncated to 1; status from first arg
	status, _, _, _ := resolveCallerArgs(fnDecl, call, info, info)
	if status != "200" {
		t.Fatalf("expected status 200 from single arg, got %q", status)
	}
}

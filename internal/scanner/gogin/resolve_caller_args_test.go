//ff:func feature=scan type=test control=sequence
//ff:what TestResolveCallerArgs_EmptyParams 테스트
package gogin

import (
	"go/ast"
	rcasTok "go/token"
	rcasPars "go/parser"
	rcasImp "go/importer"
	rcasTyp "go/types"
	"testing"
)

func TestResolveCallerArgs_EmptyParams(t *testing.T) {
	fn := &ast.FuncDecl{Type: &ast.FuncType{}}
	call := &ast.CallExpr{}
	status, tn, fields, conf := resolveCallerArgs(fn, call, nil, nil)
	if status != "" || tn != "" || fields != nil || conf != "" {
		t.Fatal("expected empty results")
	}
}


func TestResolveCallerArgs_ResponseTypeResult(t *testing.T) {
	src := `package m
type Out struct { V int ` + "`json:\"v\"`" + ` }
func write(interface{}) {}
func h() { var o Out; write(o) }
`
	fset := rcasTok.NewFileSet()
	file, _ := rcasPars.ParseFile(fset, "m.go", src, 0)
	conf := rcasTyp.Config{Importer: rcasImp.Default()}
	info := &rcasTyp.Info{
		Types: map[ast.Expr]rcasTyp.TypeAndValue{},
		Defs:  map[*ast.Ident]rcasTyp.Object{},
		Uses:  map[*ast.Ident]rcasTyp.Object{},
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
	_, tn, _, _ := resolveCallerArgs(fnDecl, call, info, info)
	if tn != "Out" {
		t.Fatalf("expected typeName Out, got %q", tn)
	}
}

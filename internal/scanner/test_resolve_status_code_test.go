//ff:func feature=scan type=extract control=sequence
//ff:what TestResolveStatusCode 리터럴 및 단순 상태 코드 해석 테스트
package scanner

import (
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"testing"
)

func TestResolveStatusCode(t *testing.T) {
	t.Run("int literal", func(t *testing.T) {
		expr := &ast.BasicLit{Kind: token.INT, Value: "200"}
		got := resolveStatusCode(expr, nil)
		if got != "200" {
			t.Errorf("expected '200', got %q", got)
		}
	})

	t.Run("nil info", func(t *testing.T) {
		expr := &ast.Ident{Name: "statusOK"}
		got := resolveStatusCode(expr, nil)
		if got != "(unknown)" {
			t.Errorf("expected '(unknown)', got %q", got)
		}
	})

	t.Run("unknown expr", func(t *testing.T) {
		expr := &ast.CompositeLit{}
		info := &types.Info{
			Types: make(map[ast.Expr]types.TypeAndValue),
			Uses:  make(map[*ast.Ident]types.Object),
		}
		got := resolveStatusCode(expr, info)
		if got != "(unknown)" {
			t.Errorf("expected '(unknown)', got %q", got)
		}
	})

	t.Run("const via SelectorExpr", func(t *testing.T) {
		fset := token.NewFileSet()
		file, err := parser.ParseFile(fset, "test.go", resolveStatusCodeConstSrc, 0)
		if err != nil {
			t.Fatal(err)
		}

		conf := types.Config{}
		info := &types.Info{
			Types: make(map[ast.Expr]types.TypeAndValue),
			Defs:  make(map[*ast.Ident]types.Object),
			Uses:  make(map[*ast.Ident]types.Object),
		}
		_, err = conf.Check("test", fset, []*ast.File{file}, info)
		if err != nil {
			t.Fatal(err)
		}

		var fnF *ast.FuncDecl
		for _, d := range file.Decls {
			fn, ok := d.(*ast.FuncDecl)
			if ok && fn.Name.Name == "f" {
				fnF = fn
				break
			}
		}
		if fnF == nil {
			t.Fatal("function f not found")
		}

		retStmt := fnF.Body.List[0].(*ast.ReturnStmt)
		ident := retStmt.Results[0].(*ast.Ident)

		got := resolveStatusCode(ident, info)
		if got != "200" {
			t.Errorf("expected '200', got %q", got)
		}
	})
}

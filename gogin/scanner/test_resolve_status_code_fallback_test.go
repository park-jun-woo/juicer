//ff:func feature=scan type=extract control=sequence
//ff:what TestResolveStatusCodeFallback TypeAndValue fallback 경로 상태 코드 해석 테스트
package scanner

import (
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"testing"
)

func TestResolveStatusCodeFallback(t *testing.T) {
	t.Run("TypeAndValue fallback", func(t *testing.T) {
		fset := token.NewFileSet()
		file, err := parser.ParseFile(fset, "test.go", resolveStatusCodeFallbackSrc, 0)
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
		retStmt := fnF.Body.List[0].(*ast.ReturnStmt)
		binExpr := retStmt.Results[0]

		got := resolveStatusCode(binExpr, info)
		if got != "404" {
			t.Errorf("expected '404', got %q", got)
		}
	})
}

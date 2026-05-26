//ff:func feature=scan type=test control=sequence
//ff:what TestScanFile_WithGinImportCov 테스트
package gogin

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestScanFile_WithGinImportCov(t *testing.T) {
	fset := token.NewFileSet()
	file := &ast.File{
		Name:    &ast.Ident{Name: "main"},
		Imports: []*ast.ImportSpec{{Path: &ast.BasicLit{Value: `"github.com/gin-gonic/gin"`}}},
		Decls: []ast.Decl{
			&ast.FuncDecl{
				Name: &ast.Ident{Name: "main"},
				Type: &ast.FuncType{},
				Body: &ast.BlockStmt{},
			},
			&ast.GenDecl{Tok: token.IMPORT},
		},
	}
	_, _ = scanFile(file, "main.go", fset)
}

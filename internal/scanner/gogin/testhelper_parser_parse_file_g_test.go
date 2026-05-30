//ff:func feature=scan type=test control=sequence
//ff:what parserParseFileG 테스트 헬퍼
package gogin

import (
	"go/ast"
	goparser "go/parser"
	"go/token"
)

func parserParseFileG(fset *token.FileSet, src string) (*ast.File, error) {
	return goparser.ParseFile(fset, "m.go", src, 0)
}

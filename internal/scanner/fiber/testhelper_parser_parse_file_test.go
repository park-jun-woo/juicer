//ff:func feature=scan type=test control=sequence
//ff:what parserParseFile 테스트 헬퍼
package fiber

import (
	"go/ast"
	"go/parser"
	"go/token"
)

func parserParseFile(fset *token.FileSet, src string) (*ast.File, error) {
	return parser.ParseFile(fset, "m.go", src, 0)
}

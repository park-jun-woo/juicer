//ff:func feature=scan type=extract control=sequence topic=spring
//ff:what AST에서 클래스 선언을 찾아 필드를 해석한다
package spring

import (
	sitter "github.com/smacker/go-tree-sitter"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func resolveClassFromAST(root *sitter.Node, src []byte, className, filePath, projectRoot string, cache map[string][]scanner.Field) []scanner.Field {
	fields, _ := resolveClassFromASTWithParams(root, src, className, filePath, projectRoot, cache)
	return fields
}

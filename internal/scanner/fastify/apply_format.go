//ff:func feature=scan type=extract control=sequence topic=fastify
//ff:what JSON Schema format 속성을 필드 타입에 적용한다
package fastify

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	sitter "github.com/smacker/go-tree-sitter"
)

func applyFormat(f *scanner.Field, propNode *sitter.Node, src []byte) {
	formatStr := extractPairStringOrIdent(propNode, src, "format")
	if formatStr != "" && f.Type == "string" {
		f.Type = formatStr
	}
}

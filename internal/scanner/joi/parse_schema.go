//ff:func feature=scan type=parse control=sequence topic=joi
//ff:what Joi.object().keys({...}) AST를 scanner.Field 슬라이스로 변환한다
package joi

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	sitter "github.com/smacker/go-tree-sitter"
)

// ParseSchema — Joi.object().keys({...}) → Field 슬라이스.
// 중첩/배열/조건부 등 미지원 형태는 빈 슬라이스로 격하한다 (크래시 금지).
func ParseSchema(node *sitter.Node, src []byte) []scanner.Field {
	if node == nil {
		return nil
	}
	objNode := FindKeysObject(node, src)
	if objNode == nil {
		return nil
	}
	return ParseObjectProperties(objNode, src)
}

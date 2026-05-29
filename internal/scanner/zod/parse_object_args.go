//ff:func feature=scan type=parse control=sequence topic=zod
//ff:what z.object() 호출의 인자에서 object 리터럴을 찾아 프로퍼티를 파싱한다
package zod

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	sitter "github.com/smacker/go-tree-sitter"
)

// ParseObjectArgs — z.object({...}) 인자 → Field 슬라이스
func ParseObjectArgs(callNode *sitter.Node, src []byte) []scanner.Field {
	args := findChildByType(callNode, "arguments")
	if args == nil {
		return nil
	}
	objNode := findChildByType(args, "object")
	if objNode == nil {
		return nil
	}
	return ParseObjectProperties(objNode, src)
}

//ff:func feature=scan type=parse control=iteration dimension=1 topic=zod
//ff:what object 리터럴의 pair 노드를 순회하여 Zod 필드를 파싱한다
package zod

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	sitter "github.com/smacker/go-tree-sitter"
)

// ParseObjectProperties — object pair 순회 → Field 슬라이스
func ParseObjectProperties(objNode *sitter.Node, src []byte) []scanner.Field {
	var fields []scanner.Field
	pairs := childrenOfType(objNode, "pair")
	for _, pair := range pairs {
		f := ParsePair(pair, src)
		if f != nil {
			fields = append(fields, *f)
		}
	}
	return fields
}

//ff:func feature=scan type=extract control=iteration dimension=1 topic=django
//ff:what Serializer 필드의 키워드 인자에서 제약조건을 추출한다
package django

import (
	sitter "github.com/smacker/go-tree-sitter"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

// applySerializerFieldConstraints applies constraints from keyword arguments.
func applySerializerFieldConstraints(field *scanner.Field, args *sitter.Node, src []byte, fieldType string) {
	for i := 0; i < int(args.ChildCount()); i++ {
		child := args.Child(i)
		if child.Type() != "keyword_argument" {
			continue
		}
		keyNode := findChildByType(child, "identifier")
		if keyNode == nil {
			continue
		}
		applyOneConstraint(field, nodeText(keyNode, src), child, src)
	}
}

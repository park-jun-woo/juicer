//ff:func feature=scan type=extract control=iteration dimension=1 topic=actix
//ff:what #[serde(...)] 어트리뷰트를 파싱한다
package actix

import (
	"strings"

	sitter "github.com/smacker/go-tree-sitter"
)

func parseSerdeAttribute(attrItem *sitter.Node, src []byte) *serdeAttr {
	tokenTree := serdeTokenTree(attrItem, src)
	if tokenTree == nil {
		return nil
	}

	result := &serdeAttr{}
	text := nodeText(tokenTree, src)
	text = strings.TrimPrefix(text, "(")
	text = strings.TrimSuffix(text, ")")

	for _, part := range strings.Split(text, ",") {
		applySerdePart(result, strings.TrimSpace(part))
	}

	return result
}

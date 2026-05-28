//ff:func feature=scan type=extract control=sequence topic=actix
//ff:what #[serde(...)] 어트리뷰트를 파싱한다
package actix

import (
	"strings"

	sitter "github.com/smacker/go-tree-sitter"
)

type serdeAttr struct {
	rename     string
	hasDefault bool
	skip       bool
}

func parseSerdeAttribute(attrItem *sitter.Node, src []byte) *serdeAttr {
	attr := findChildByType(attrItem, "attribute")
	if attr == nil {
		return nil
	}
	nameNode := findChildByType(attr, "identifier")
	if nameNode == nil {
		return nil
	}
	if nodeText(nameNode, src) != "serde" {
		return nil
	}

	tokenTree := findChildByType(attr, "token_tree")
	if tokenTree == nil {
		return nil
	}

	result := &serdeAttr{}
	text := nodeText(tokenTree, src)
	// Remove surrounding parentheses
	text = strings.TrimPrefix(text, "(")
	text = strings.TrimSuffix(text, ")")

	parts := strings.Split(text, ",")
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "default" {
			result.hasDefault = true
		}
		if part == "skip" || part == "skip_deserializing" {
			result.skip = true
		}
		if strings.HasPrefix(part, "rename") {
			eqIdx := strings.Index(part, "=")
			if eqIdx >= 0 {
				val := strings.TrimSpace(part[eqIdx+1:])
				result.rename = unquoteRust(val)
			}
		}
	}

	return result
}

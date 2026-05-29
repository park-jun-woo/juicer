//ff:func feature=scan type=extract control=selection topic=actix
//ff:what field_declaration_list 자식 노드 하나를 처리해 필드/대기 어트리뷰트 상태를 갱신한다
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func consumeFieldChild(child *sitter.Node, src []byte, fields []scanner.Field, pending []serdeAttr) ([]scanner.Field, []serdeAttr) {
	switch child.Type() {
	case "attribute_item":
		attr := parseSerdeAttribute(child, src)
		if attr != nil {
			pending = append(pending, *attr)
		}
		return fields, pending
	case "field_declaration":
		f := parseOneField(child, src, pending)
		if f != nil {
			fields = append(fields, *f)
		}
		return fields, nil
	case ",":
		return fields, pending
	default:
		return fields, nil
	}
}

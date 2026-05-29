//ff:func feature=scan type=extract control=iteration dimension=1 topic=actix
//ff:what Rust struct의 필드를 파싱한다
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func extractStructFields(structNode *sitter.Node, src []byte) []scanner.Field {
	fieldList := findChildByType(structNode, "field_declaration_list")
	if fieldList == nil {
		return nil
	}

	var fields []scanner.Field
	var pendingAttrs []serdeAttr

	for i := 0; i < int(fieldList.ChildCount()); i++ {
		child := fieldList.Child(i)
		fields, pendingAttrs = consumeFieldChild(child, src, fields, pendingAttrs)
	}

	return fields
}

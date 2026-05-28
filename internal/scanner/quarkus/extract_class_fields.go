//ff:func feature=scan type=extract control=iteration dimension=1 topic=quarkus
//ff:what DTO/Request 클래스 필드를 파싱한다
package quarkus

import (
	sitter "github.com/smacker/go-tree-sitter"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func extractClassFields(cls *sitter.Node, src []byte) []scanner.Field {
	body := findChildByType(cls, "class_body")
	if body == nil {
		return nil
	}
	var fields []scanner.Field
	for i := 0; i < int(body.ChildCount()); i++ {
		child := body.Child(i)
		if child.Type() != "field_declaration" {
			continue
		}
		if isStaticField(child) {
			continue
		}
		f := extractOneField(child, src)
		if f.Name != "" {
			fields = append(fields, f)
		}
	}
	return fields
}

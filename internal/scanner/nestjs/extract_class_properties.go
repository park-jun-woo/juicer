//ff:func feature=scan type=extract control=iteration dimension=1 topic=nestjs
//ff:what 클래스 본문에서 프로퍼티 정의를 추출한다
package nestjs

import sitter "github.com/smacker/go-tree-sitter"

// extractClassProperties extracts property definitions from a class body.
func extractClassProperties(cls *sitter.Node, src []byte) []dtoField {
	body := findChildByType(cls, "class_body")
	if body == nil {
		return nil
	}
	props := childrenOfType(body, "public_field_definition")
	var fields []dtoField
	for _, prop := range props {
		f := extractOneProperty(prop, src)
		if f.name != "" {
			fields = append(fields, f)
		}
	}
	return fields
}

//ff:func feature=scan type=extract control=iteration dimension=1 topic=nestjs
//ff:what 단일 프로퍼티 정의에서 필드 정보를 추출한다
package nestjs

import sitter "github.com/smacker/go-tree-sitter"

// extractOneProperty extracts info from a single public_field_definition.
func extractOneProperty(prop *sitter.Node, src []byte) dtoField {
	var f dtoField
	decorators := propertyDecorators(prop, src)
	for _, d := range decorators {
		f.validators = append(f.validators, d.name)
		if d.name == "IsOptional" {
			f.optional = true
		}
	}
	for i := 0; i < int(prop.ChildCount()); i++ {
		child := prop.Child(i)
		switch child.Type() {
		case "property_identifier":
			f.name = nodeText(child, src)
		case "type_annotation":
			f.tsType = extractTypeAnnotation(child, src)
		case "?":
			f.optional = true
		}
	}
	return f
}

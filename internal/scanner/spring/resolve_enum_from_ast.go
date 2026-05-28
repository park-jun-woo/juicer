//ff:func feature=scan type=extract control=iteration dimension=1 topic=spring
//ff:what AST에서 enum 선언을 찾아 필드로 변환한다
package spring

import (
	sitter "github.com/smacker/go-tree-sitter"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func resolveEnumFromAST(root *sitter.Node, src []byte, className string, cache map[string][]scanner.Field) []scanner.Field {
	enums := findAllByType(root, "enum_declaration")
	for _, en := range enums {
		nameNode := findChildByType(en, "identifier")
		if nameNode == nil || nodeText(nameNode, src) != className {
			continue
		}
		values := extractEnumValues(en, src)
		field := scanner.Field{
			Name: className,
			Type: "string",
			Enum: values,
		}
		result := []scanner.Field{field}
		cache[className] = result
		return result
	}
	return nil
}

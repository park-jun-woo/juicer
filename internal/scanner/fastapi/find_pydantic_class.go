//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastapi
//ff:what 지정 이름의 Pydantic 클래스를 찾아 필드를 반환한다
package fastapi

import (
	sitter "github.com/smacker/go-tree-sitter"

	"github.com/park-jun-woo/juicer/internal/scanner"
)

// findPydanticClass finds a class with the given name and extracts its fields.
func findPydanticClass(root *sitter.Node, src []byte, className string) []scanner.Field {
	classes := findAllByType(root, "class_definition")
	for _, cls := range classes {
		nameNode := findChildByType(cls, "identifier")
		if nameNode == nil {
			continue
		}
		if nodeText(nameNode, src) != className {
			continue
		}
		if !isBaseModelSubclass(cls, src) {
			continue
		}
		fields := extractPydanticFields(cls, src)
		return pydanticFieldsToScannerFields(fields)
	}
	return nil
}

//ff:func feature=scan type=extract control=sequence topic=django
//ff:what Serializer 필드 호출에서 필드 타입 이름을 추출한다
package django

import (
	"strings"

	sitter "github.com/smacker/go-tree-sitter"
)

// extractDRFFieldType extracts the field type from a serializer field call.
// e.g., serializers.CharField() -> "CharField", CharField() -> "CharField"
func extractDRFFieldType(callNode *sitter.Node, src []byte) string {
	attrNode := findChildByType(callNode, "attribute")
	if attrNode != nil {
		text := nodeText(attrNode, src)
		parts := strings.Split(text, ".")
		return parts[len(parts)-1]
	}
	identNode := findChildByType(callNode, "identifier")
	if identNode != nil {
		return nodeText(identNode, src)
	}
	return ""
}

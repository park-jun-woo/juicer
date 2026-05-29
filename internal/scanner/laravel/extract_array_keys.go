//ff:func feature=scan type=extract control=iteration dimension=1 topic=laravel
//ff:what toArray() 반환 배열의 키들을 string 타입 필드로 추출한다
package laravel

import (
	sitter "github.com/smacker/go-tree-sitter"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

// extractArrayKeys extracts keys from the return array of toArray().
func extractArrayKeys(method *sitter.Node, src []byte) []scanner.Field {
	arr := methodReturnArray(method)
	if arr == nil {
		return nil
	}
	var fields []scanner.Field
	for _, elem := range childrenOfType(arr, "array_element_initializer") {
		keyNode := findChildByType(elem, "string")
		if keyNode == nil {
			continue
		}
		key := extractStringContent(keyNode, src)
		if key == "" {
			continue
		}
		fields = append(fields, scanner.Field{
			Name: key,
			JSON: key,
			Type: "string", // default since we can't infer type from $this->field
		})
	}
	return fields
}

//ff:func feature=scan type=extract control=iteration dimension=1 topic=laravel
//ff:what rules() 메서드 반환 배열에서 필드들을 추출한다
package laravel

import (
	sitter "github.com/smacker/go-tree-sitter"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

// extractFieldsFromRulesMethod extracts fields from the return array of a rules() method.
func extractFieldsFromRulesMethod(method *sitter.Node, src []byte) []scanner.Field {
	arr := methodReturnArray(method)
	if arr == nil {
		return nil
	}
	var fields []scanner.Field
	for _, elem := range childrenOfType(arr, "array_element_initializer") {
		field := extractOneRuleField(elem, src)
		if field != nil {
			fields = append(fields, *field)
		}
	}
	return fields
}

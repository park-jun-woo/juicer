//ff:func feature=scan type=extract control=sequence topic=laravel
//ff:what 'name' => 'rules' 또는 'name' => ['rules'] 항목에서 필드 하나를 추출한다
package laravel

import (
	sitter "github.com/smacker/go-tree-sitter"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

// extractOneRuleField extracts a single field from 'name' => 'rules' or 'name' => ['rules'].
func extractOneRuleField(elem *sitter.Node, src []byte) *scanner.Field {
	if elem.ChildCount() < 3 {
		return nil
	}
	keyNode := findChildByType(elem, "string")
	if keyNode == nil {
		return nil
	}
	fieldName := extractStringContent(keyNode, src)
	if fieldName == "" {
		return nil
	}

	rules := extractRuleStrings(elem, src)
	if len(rules) == 0 {
		return nil
	}

	field := laravelRulesToField(fieldName, rules)
	return &field
}

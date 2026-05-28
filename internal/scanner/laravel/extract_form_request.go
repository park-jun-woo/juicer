//ff:func feature=scan type=extract control=iteration dimension=1 topic=laravel
//ff:what FormRequest::rules() 배열에서 필드명과 유효성 규칙을 추출한다
package laravel

import (
	"strings"

	sitter "github.com/smacker/go-tree-sitter"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

// extractFormRequest finds a FormRequest class and extracts fields from its rules() method.
func extractFormRequest(absRoot, className string, parsedFiles map[string]*fileInfo) []scanner.Field {
	fi := findFormRequestFile(absRoot, className, parsedFiles)
	if fi == nil {
		return nil
	}
	return extractRulesFromFile(fi, className)
}

// findFormRequestFile locates the file containing the FormRequest class.
func findFormRequestFile(absRoot, className string, parsedFiles map[string]*fileInfo) *fileInfo {
	// Search in parsed files first
	for _, fi := range parsedFiles {
		if classMatches(fi, className) {
			return fi
		}
	}
	// Try PSR-4 convention
	candidates := []string{
		absRoot + "/app/Http/Requests/" + className + ".php",
	}
	for _, candidate := range candidates {
		fi, err := parseFile(absRoot, candidate)
		if err == nil {
			return fi
		}
	}
	return nil
}

// extractRulesFromFile finds the rules() method in the class and extracts field definitions.
func extractRulesFromFile(fi *fileInfo, className string) []scanner.Field {
	classes := findAllByType(fi.root, "class_declaration")
	for _, cls := range classes {
		nameNode := findChildByType(cls, "name")
		if nameNode != nil && nodeText(nameNode, fi.src) != className {
			continue
		}
		declList := findChildByType(cls, "declaration_list")
		if declList == nil {
			continue
		}
		methods := childrenOfType(declList, "method_declaration")
		for _, method := range methods {
			mName := findChildByType(method, "name")
			if mName == nil || nodeText(mName, fi.src) != "rules" {
				continue
			}
			return extractFieldsFromRulesMethod(method, fi.src)
		}
	}
	return nil
}

// extractFieldsFromRulesMethod extracts fields from the return array of a rules() method.
func extractFieldsFromRulesMethod(method *sitter.Node, src []byte) []scanner.Field {
	retStmts := findAllByType(method, "return_statement")
	if len(retStmts) == 0 {
		return nil
	}
	arrNodes := findAllByType(retStmts[0], "array_creation_expression")
	if len(arrNodes) == 0 {
		return nil
	}
	arr := arrNodes[0]
	elems := childrenOfType(arr, "array_element_initializer")
	var fields []scanner.Field
	for _, elem := range elems {
		field := extractOneRuleField(elem, src)
		if field != nil {
			fields = append(fields, *field)
		}
	}
	return fields
}

// extractOneRuleField extracts a single field from 'name' => 'rules' or 'name' => ['rules'].
func extractOneRuleField(elem *sitter.Node, src []byte) *scanner.Field {
	// Must be a key => value pair
	if elem.ChildCount() < 3 {
		return nil
	}
	// Find key (string before =>)
	keyNode := findChildByType(elem, "string")
	if keyNode == nil {
		return nil
	}
	fieldName := extractStringContent(keyNode, src)
	if fieldName == "" {
		return nil
	}

	// Find value (after =>)
	rules := extractRuleStrings(elem, src)
	if len(rules) == 0 {
		return nil
	}

	field := laravelRulesToField(fieldName, rules)
	return &field
}

// extractRuleStrings extracts rule strings from the value side of a rules entry.
// Supports both pipe-delimited strings ('required|string|max:255')
// and array syntax (['required', 'string', 'max:255']).
func extractRuleStrings(elem *sitter.Node, src []byte) []string {
	// Check for array syntax first
	arr := findChildByType(elem, "array_creation_expression")
	if arr != nil {
		return extractStringArray(arr, src)
	}

	// Find the value string (the one after =>)
	strNodes := childrenOfType(elem, "string")
	if len(strNodes) < 2 {
		return nil
	}
	// Second string is the value (first is the key)
	valueStr := extractStringContent(strNodes[1], src)
	if valueStr == "" {
		return nil
	}
	return strings.Split(valueStr, "|")
}

//ff:func feature=scan type=extract control=sequence topic=supafunc
//ff:what const body = await req.json() 후 body.field 멤버 접근에서 필드명을 추출한다
package supafunc

import (
	"strings"

	sitter "github.com/smacker/go-tree-sitter"
)

func extractBodyMemberAccess(body *sitter.Node, src []byte) []string {
	// 1. Find variable name assigned from req.json()
	var varName string
	walkNodes(body, func(n *sitter.Node) {
		if varName != "" {
			return
		}
		if n.Type() != "lexical_declaration" && n.Type() != "variable_declaration" {
			return
		}
		text := nodeText(n, src)
		if !strings.Contains(text, "req.json()") {
			return
		}
		declarators := childrenOfType(n, "variable_declarator")
		for _, decl := range declarators {
			nameNode := findChildByType(decl, "identifier")
			if nameNode == nil {
				continue
			}
			// Skip if it's an object_pattern (destructuring)
			if findChildByType(decl, "object_pattern") != nil {
				continue
			}
			varName = nodeText(nameNode, src)
		}
	})
	if varName == "" {
		return nil
	}

	// 2. Collect member_expression where object == varName (first-level only)
	seen := map[string]bool{}
	var fields []string
	walkNodes(body, func(n *sitter.Node) {
		if n.Type() != "member_expression" {
			return
		}
		// Only first-level: object must be an identifier, not another member_expression
		obj := findChildByType(n, "identifier")
		if obj == nil {
			return
		}
		if nodeText(obj, src) != varName {
			return
		}
		// Ensure the identifier is the object (left side), not the property
		prop := findChildByType(n, "property_identifier")
		if prop == nil {
			return
		}
		name := nodeText(prop, src)
		if name != "" && !seen[name] {
			seen[name] = true
			fields = append(fields, name)
		}
	})
	return fields
}

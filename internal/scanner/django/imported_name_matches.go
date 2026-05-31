//ff:func feature=scan type=extract control=selection topic=django
//ff:what import_from_statement 자식 노드가 지정 이름의 import인지 판별한다
package django

import sitter "github.com/smacker/go-tree-sitter"

// importedNameMatches reports whether a single child of an import_from_statement
// names the given import (directly as a dotted_name or via aliased_import).
func importedNameMatches(child *sitter.Node, name string, src []byte) bool {
	switch child.Type() {
	case "dotted_name":
		return nodeText(child, src) == name
	case "aliased_import":
		dn := findChildByType(child, "dotted_name")
		return dn != nil && nodeText(dn, src) == name
	default:
		return false
	}
}

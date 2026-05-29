//ff:func feature=scan type=extract control=sequence topic=flask
//ff:what aliased_import 노드에서 (로컬명, 원본명)을 파싱한다
package flask

import sitter "github.com/smacker/go-tree-sitter"

// parseAliasedImport extracts (localName, originalName) from an aliased_import node.
// e.g., "auth as auth_blueprint" -> ("auth_blueprint", "auth").
// Returns ("", "") if the node shape is unexpected.
func parseAliasedImport(node *sitter.Node, src []byte) (string, string) {
	origNode := findChildByType(node, "dotted_name")
	localNode := findChildByType(node, "identifier")
	if origNode == nil || localNode == nil {
		return "", ""
	}
	return nodeText(localNode, src), nodeText(origNode, src)
}

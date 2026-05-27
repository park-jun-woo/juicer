//ff:func feature=scan type=extract control=selection topic=fastapi
//ff:what 자식 노드가 import된 이름인지 확인하여 반환한다
package fastapi

import sitter "github.com/smacker/go-tree-sitter"

// tryImportedName returns the imported name if this child is an imported identifier.
func tryImportedName(child *sitter.Node, stmt *sitter.Node, idx int, src []byte) string {
	if idx == 0 {
		return ""
	}
	prev := stmt.Child(idx - 1)
	prevText := nodeText(prev, src)

	switch child.Type() {
	case "dotted_name":
		if prevText == "import" || prevText == "," {
			return nodeText(child, src)
		}
	case "identifier":
		if prevText == "import" || prevText == "," {
			return nodeText(child, src)
		}
	}
	return ""
}

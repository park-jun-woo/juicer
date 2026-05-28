//ff:func feature=scan type=extract control=iteration dimension=1 topic=django
//ff:what 리스트 노드의 자식 call 노드에서 path() 호출을 파싱한다
package django

import sitter "github.com/smacker/go-tree-sitter"

// parsePathCallsInList parses all path() calls in a list node.
func parsePathCallsInList(listNode *sitter.Node, src []byte) []urlEntry {
	var entries []urlEntry
	for _, child := range childrenOfType(listNode, "call") {
		entry := parsePathCall(child, src)
		if entry != nil {
			entries = append(entries, *entry)
		}
	}
	return entries
}

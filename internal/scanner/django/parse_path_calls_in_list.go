//ff:func feature=scan type=extract control=iteration dimension=1 topic=django
//ff:what 노드의 직접 자식 call에서 path()/re_path() 호출을 파싱한다
package django

import sitter "github.com/smacker/go-tree-sitter"

// parsePathCallsInList parses all path()/re_path() calls among a node's direct call children
// (used for both list literals and wrapper argument lists like i18n_patterns(...)).
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

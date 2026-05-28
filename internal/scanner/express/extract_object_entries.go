//ff:func feature=scan type=extract control=iteration dimension=1 topic=express
//ff:what 배열 노드의 각 object에서 path(string)와 route(identifier) 값을 추출한다
package express

import sitter "github.com/smacker/go-tree-sitter"

func extractObjectEntries(arr *sitter.Node, src []byte) []pathRouteEntry {
	var entries []pathRouteEntry
	for _, obj := range childrenOfType(arr, "object") {
		e := extractPathRouteFromObject(obj, src)
		if e != nil {
			entries = append(entries, *e)
		}
	}
	return entries
}

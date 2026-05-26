//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastapi
//ff:what FastAPI/APIRouter 인스턴스와 include_router 호출에서 접두사 체인을 해석한다
package fastapi

import sitter "github.com/smacker/go-tree-sitter"

// resolveRouterPrefixes finds all router variable assignments and builds
// a map of variable name -> full prefix (including include_router chains).
func resolveRouterPrefixes(root *sitter.Node, src []byte) map[string]string {
	routers := findRouterAssignments(root, src)
	includes := findIncludeRouterCalls(root, src)

	prefixes := make(map[string]string)
	for _, r := range routers {
		prefixes[r.varName] = r.prefix
	}
	for _, inc := range includes {
		parentPrefix := prefixes[inc.parentVar]
		childPrefix := prefixes[inc.childVar]
		prefixes[inc.childVar] = joinPath(parentPrefix, inc.extraPrefix, childPrefix)
	}
	return prefixes
}

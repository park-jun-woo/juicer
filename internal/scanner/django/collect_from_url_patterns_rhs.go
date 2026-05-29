//ff:func feature=scan type=extract control=sequence topic=django
//ff:what urlpatterns RHS가 list든 i18n_patterns(...) 같은 래퍼 call이든 path() 호출을 수집한다
package django

import sitter "github.com/smacker/go-tree-sitter"

// collectFromURLPatternsRHS collects path() calls from a urlpatterns assignment RHS,
// handling both `urlpatterns = [...]` (list) and `urlpatterns = i18n_patterns(...)` (call wrapper).
func collectFromURLPatternsRHS(assignNode *sitter.Node, src []byte) []urlEntry {
	if listNode := findChildByType(assignNode, "list"); listNode != nil {
		return parsePathCallsInList(listNode, src)
	}
	if callNode := findChildByType(assignNode, "call"); callNode != nil {
		if argList := findChildByType(callNode, "argument_list"); argList != nil {
			return parsePathCallsInList(argList, src)
		}
	}
	return nil
}

//ff:func feature=scan type=extract control=iteration dimension=1 topic=django
//ff:what as_view({...}) 호출 인자의 method→action dict를 파싱한다
package django

import (
	"strings"

	sitter "github.com/smacker/go-tree-sitter"
)

// parseAsViewDict reads the {"get": "list", ...} dict from an .as_view(...) call
// node's argument list and returns the lowercase HTTP-method -> action map.
// Returns nil when the call has no dict literal argument.
func parseAsViewDict(callNode *sitter.Node, src []byte) map[string]string {
	args := findChildByType(callNode, "argument_list")
	if args == nil {
		return nil
	}
	dict := findChildByType(args, "dictionary")
	if dict == nil {
		return nil
	}
	result := map[string]string{}
	for _, pair := range childrenOfType(dict, "pair") {
		strs := childrenOfType(pair, "string")
		if len(strs) < 2 {
			continue
		}
		method := strings.ToLower(unquotePython(nodeText(strs[0], src)))
		action := unquotePython(nodeText(strs[1], src))
		if method != "" && action != "" {
			result[method] = action
		}
	}
	if len(result) == 0 {
		return nil
	}
	return result
}

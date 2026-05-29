//ff:func feature=scan type=extract control=sequence topic=laravel
//ff:what member_call이 json() 호출이면 두 번째 인자(상태 코드) 텍스트를 반환한다
package laravel

import (
	"strings"

	sitter "github.com/smacker/go-tree-sitter"
)

func jsonCallStatusCode(mc *sitter.Node, src []byte) string {
	if lastMemberCallName(mc, src) != "json" {
		return ""
	}
	jsonArgs := findChildByType(mc, "arguments")
	if jsonArgs == nil {
		return ""
	}
	argList := childrenOfType(jsonArgs, "argument")
	if len(argList) < 2 {
		return ""
	}
	return strings.TrimSpace(nodeText(argList[1], src))
}

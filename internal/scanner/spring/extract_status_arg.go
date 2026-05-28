//ff:func feature=scan type=extract control=iteration dimension=1 topic=spring
//ff:what ResponseEntity.status() 호출의 인자에서 상태 코드를 추출한다
package spring

import sitter "github.com/smacker/go-tree-sitter"

func extractStatusArg(inv *sitter.Node, src []byte) string {
	for _, argList := range findAllByType(inv, "argument_list") {
		code := extractStatusFromArgList(argList, src)
		if code != "" {
			return code
		}
	}
	return ""
}

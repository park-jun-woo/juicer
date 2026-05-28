//ff:func feature=scan type=extract control=iteration dimension=1 topic=quarkus
//ff:what Response.status() 호출의 인자에서 상태 코드를 추출한다
package quarkus

import sitter "github.com/smacker/go-tree-sitter"

func extractResponseStatusArg(inv *sitter.Node, src []byte) string {
	for _, argList := range findAllByType(inv, "argument_list") {
		code := extractIntLiteralFromArgList(argList, src)
		if code != "" {
			return code
		}
	}
	return ""
}

//ff:func feature=scan type=extract control=iteration dimension=1 topic=laravel
//ff:what json() 호출의 두 번째 인자에서 상태 코드를 추출한다(기본 "200")
package laravel

import (
	sitter "github.com/smacker/go-tree-sitter"
)

func jsonResponseStatus(retNode *sitter.Node, src []byte) string {
	status := "200"
	for _, mc := range findAllByType(retNode, "member_call_expression") {
		if code := jsonCallStatusCode(mc, src); code != "" {
			status = code
		}
	}
	return status
}

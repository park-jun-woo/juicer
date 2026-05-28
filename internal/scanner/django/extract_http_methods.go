//ff:func feature=scan type=extract control=iteration dimension=1 topic=django
//ff:what APIView 클래스 body에서 HTTP 메서드 정의를 추출한다
package django

import (
	"strings"

	sitter "github.com/smacker/go-tree-sitter"
)

// extractHTTPMethods finds HTTP method definitions (get, post, etc.) in a class body.
func extractHTTPMethods(body *sitter.Node, src []byte) []string {
	var methods []string
	for _, funcDef := range childrenOfType(body, "function_definition") {
		nameNode := findChildByType(funcDef, "identifier")
		if nameNode == nil {
			continue
		}
		name := nodeText(nameNode, src)
		if method, ok := apiviewHTTPMethods[strings.ToLower(name)]; ok {
			methods = append(methods, method)
		}
	}
	return methods
}

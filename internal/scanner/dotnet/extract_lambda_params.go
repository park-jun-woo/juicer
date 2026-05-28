//ff:func feature=scan type=extract control=iteration dimension=1 topic=dotnet
//ff:what 람다 파라미터에서 요청 정보를 추출한다
package dotnet

import (
	sitter "github.com/smacker/go-tree-sitter"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func extractLambdaParams(lambda *sitter.Node, src []byte, path string) *scanner.Request {
	params := findChildByType(lambda, "parameter_list")
	if params == nil {
		return nil
	}
	var req scanner.Request
	for _, param := range childrenOfType(params, "parameter") {
		classifyLambdaParam(param, src, path, &req)
	}
	if hasContent(&req) {
		return &req
	}
	return nil
}

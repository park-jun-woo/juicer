//ff:func feature=scan type=extract control=sequence topic=dotnet
//ff:what argument 목록에서 람다 핸들러의 요청 정보를 추출한다
package dotnet

import (
	sitter "github.com/smacker/go-tree-sitter"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func extractLambdaRequest(argNodes []*sitter.Node, src []byte, path string) *scanner.Request {
	if len(argNodes) < 2 {
		return nil
	}
	lambda := findChildByType(argNodes[1], "lambda_expression")
	if lambda == nil {
		return nil
	}
	return extractLambdaParams(lambda, src, path)
}

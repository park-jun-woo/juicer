//ff:func feature=scan type=extract control=sequence topic=supafunc
//ff:what 특정 메서드의 블록 노드에서 request와 response를 추출하여 Endpoint를 구성한다
package supafunc

import (
	sitter "github.com/smacker/go-tree-sitter"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func extractEndpointForMethod(block *sitter.Node, src []byte, method, path, handler, file string) scanner.Endpoint {
	req := buildRequest(
		extractRequestJSON(block, src),
		extractQueryParams(block, src),
	)
	responses := buildResponses(extractResponseStatus(block, src))
	return scanner.Endpoint{
		Method:    method,
		Path:      path,
		Handler:   handler,
		File:      file,
		Request:   req,
		Responses: responses,
	}
}

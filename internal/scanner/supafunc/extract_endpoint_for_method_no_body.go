//ff:func feature=scan type=extract control=sequence topic=supafunc
//ff:what body 없이 query params와 response만 추출하여 Endpoint를 구성한다 (폴백 body 차단용)
package supafunc

import (
	sitter "github.com/smacker/go-tree-sitter"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func extractEndpointForMethodNoBody(block *sitter.Node, src []byte, method, path, handler, file string) scanner.Endpoint {
	req := buildRequest(
		nil,
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

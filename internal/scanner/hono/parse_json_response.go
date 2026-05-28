//ff:func feature=scan type=extract control=sequence topic=hono
//ff:what c.json(body, status) 호출에서 상태 코드를 추출하여 Response를 생성한다
package hono

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	sitter "github.com/smacker/go-tree-sitter"
)

func parseJsonResponse(call *sitter.Node, src []byte) *scanner.Response {
	status := extractSecondNumberArg(call, src, "200")
	return &scanner.Response{Status: status, Kind: "json"}
}

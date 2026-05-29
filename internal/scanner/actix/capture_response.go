//ff:func feature=scan type=extract control=sequence topic=actix
//ff:what HttpResponse::<status> 노드 하나를 해석해 응답 목록에 추가한다
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func captureResponse(n *sitter.Node, src []byte, seen map[string]bool, responses *[]scanner.Response) {
	if n.Type() != "scoped_identifier" {
		return
	}
	parts := splitScoped(nodeText(n, src))
	if len(parts) != 2 || parts[0] != "HttpResponse" {
		return
	}
	code, ok := httpResponseStatuses[parts[1]]
	if !ok || seen[code] {
		return
	}
	seen[code] = true
	*responses = append(*responses, scanner.Response{
		Status: code,
		Kind:   detectResponseKind(n, src),
	})
}

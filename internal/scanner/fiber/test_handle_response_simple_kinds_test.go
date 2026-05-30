//ff:func feature=scan type=test control=iteration dimension=1
//ff:what TestHandleResponse_SimpleKinds 테스트
package fiber

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"go/ast"
	"testing"
)

func TestHandleResponse_SimpleKinds(t *testing.T) {
	for _, kind := range []string{"string", "data", "file"} {
		ep := scanner.Endpoint{}
		handleResponse(&ep, &ast.CallExpr{}, kind, nil, "handler")
		if len(ep.Responses) != 1 || ep.Responses[0].Status != "200" || ep.Responses[0].Kind != kind {
			t.Fatalf("%s: %v", kind, ep.Responses)
		}
	}
}

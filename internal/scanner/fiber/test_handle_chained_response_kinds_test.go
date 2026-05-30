//ff:func feature=scan type=test control=iteration dimension=1
//ff:what TestHandleChainedResponse_Kinds 테스트
package fiber

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"go/ast"
	"go/token"
	"testing"
)

func TestHandleChainedResponse_Kinds(t *testing.T) {
	statusArg := []ast.Expr{&ast.BasicLit{Kind: token.INT, Value: "200"}}

	for _, tc := range []struct{ method, kind string }{
		{"SendString", "string"},
		{"Send", "data"},
		{"XML", "XML"},
	} {
		ep := scanner.Endpoint{}
		sc, oc := chainCall(tc.method, statusArg)
		handleChainedResponse(&ep, sc, oc, tc.method, nil, "handler")
		if len(ep.Responses) != 1 || ep.Responses[0].Kind != tc.kind {
			t.Fatalf("%s: kind = %v", tc.method, ep.Responses)
		}
	}
}

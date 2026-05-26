//ff:func feature=scan type=test control=sequence
//ff:what TestHandleResponse_UnknownStatusCov 테스트
package gogin

import (
	"go/ast"
	"go/types"
	"testing"
	"github.com/park-jun-woo/juicer/internal/scanner"
)

func TestHandleResponse_UnknownStatusCov(t *testing.T) {
	ep := &scanner.Endpoint{}
	info := &types.Info{Uses: make(map[*ast.Ident]types.Object), Types: make(map[ast.Expr]types.TypeAndValue)}
	handleResponse(ep, &ast.CallExpr{}, "json", info, "handler")
	if ep.Responses[0].Status != "(unknown)" {
		t.Fatalf("expected (unknown), got %s", ep.Responses[0].Status)
	}
}

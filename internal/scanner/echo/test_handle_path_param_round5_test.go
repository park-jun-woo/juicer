//ff:func feature=scan type=test control=sequence topic=echo
//ff:what TestHandlePathParam_Round5 테스트
package echo

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"go/ast"
	"testing"
)

func TestHandlePathParam_Round5(t *testing.T) {
	ep := &scanner.Endpoint{}
	call := callWithStringArg(t, `"id"`)
	handlePathParam(ep, call)
	if ep.Request == nil || len(ep.Request.PathParams) != 1 || ep.Request.PathParams[0].Name != "id" {
		t.Fatalf("path params: %+v", ep.Request)
	}

	handlePathParam(ep, call)
	if len(ep.Request.PathParams) != 1 {
		t.Fatalf("duplicate added: %+v", ep.Request.PathParams)
	}

	ep2 := &scanner.Endpoint{}
	handlePathParam(ep2, &ast.CallExpr{})
	if ep2.Request != nil {
		t.Fatalf("expected nil request for no args")
	}
}

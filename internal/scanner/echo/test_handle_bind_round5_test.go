//ff:func feature=scan type=test control=sequence topic=echo
//ff:what TestHandleBind_Round5 테스트
package echo

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestHandleBind_Round5(t *testing.T) {
	ep := &scanner.Endpoint{}
	call := callExprFrom(t, `c.Bind(&dto)`)
	handleBind(ep, call, "Bind", nil)
	if ep.Request == nil || ep.Request.Body == nil || ep.Request.Body.VarName != "dto" {
		t.Fatalf("bind: %+v", ep.Request)
	}

	handleBind(ep, callExprFrom(t, `c.Bind(&other)`), "Bind", nil)
	if ep.Request.Body.VarName != "dto" {
		t.Fatalf("second bind overwrote: %+v", ep.Request.Body)
	}
}

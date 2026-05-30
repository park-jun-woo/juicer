//ff:func feature=scan type=test control=sequence
//ff:what handleBodyParser — BodyParser 바인딩 처리 테스트
package fiber

import (
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestHandleBodyParser_FirstBinding(t *testing.T) {
	ep := &scanner.Endpoint{}
	call := parseCall(t, "c.BodyParser(&req)")
	handleBodyParser(ep, call, "json", newEmptyInfo())
	if ep.Request == nil || ep.Request.Body == nil {
		t.Fatal("expected body set")
	}
	if ep.Request.Body.VarName != "req" || ep.Request.Body.Method != "json" {
		t.Fatalf("body = %+v", ep.Request.Body)
	}
}

func TestHandleBodyParser_AlreadySet(t *testing.T) {
	ep := &scanner.Endpoint{}
	scanner.EnsureRequest(ep)
	ep.Request.Body = &scanner.Body{VarName: "existing"}
	call := parseCall(t, "c.BodyParser(&other)")
	handleBodyParser(ep, call, "json", newEmptyInfo())
	if ep.Request.Body.VarName != "existing" {
		t.Fatalf("should keep first binding, got %q", ep.Request.Body.VarName)
	}
}

func TestHandleBodyParser_NoArgs(t *testing.T) {
	ep := &scanner.Endpoint{}
	call := parseCall(t, "c.BodyParser()")
	handleBodyParser(ep, call, "json", newEmptyInfo())
	if ep.Request.Body == nil || ep.Request.Body.VarName != "(unknown)" {
		t.Fatalf("expected (unknown) varname, got %+v", ep.Request.Body)
	}
}

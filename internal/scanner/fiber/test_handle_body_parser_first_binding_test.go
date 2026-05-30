//ff:func feature=scan type=test control=sequence
//ff:what TestHandleBodyParser_FirstBinding 테스트
package fiber

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
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

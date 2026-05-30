//ff:func feature=scan type=test control=sequence
//ff:what TestHandleBodyParser_AlreadySet 테스트
package fiber

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

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

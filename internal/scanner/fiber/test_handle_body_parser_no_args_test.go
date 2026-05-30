//ff:func feature=scan type=test control=sequence
//ff:what TestHandleBodyParser_NoArgs 테스트
package fiber

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestHandleBodyParser_NoArgs(t *testing.T) {
	ep := &scanner.Endpoint{}
	call := parseCall(t, "c.BodyParser()")
	handleBodyParser(ep, call, "json", newEmptyInfo())
	if ep.Request.Body == nil || ep.Request.Body.VarName != "(unknown)" {
		t.Fatalf("expected (unknown) varname, got %+v", ep.Request.Body)
	}
}

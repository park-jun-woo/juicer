//ff:func feature=scan type=test control=sequence topic=echo
//ff:what TestHandleResponse_Round5 테스트
package echo

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestHandleResponse_Round5(t *testing.T) {
	_, info := checkSrc(t, `package m
type UserDto struct { Name string `+"`json:\"name\"`"+` }
var u UserDto
var _ = u
`)
	call := callExprFrom(t, `c.JSON(200, u)`)
	ep := &scanner.Endpoint{}
	handleResponse(ep, call, "json", info, "handler")
	if len(ep.Responses) != 1 {
		t.Fatalf("responses: %+v", ep.Responses)
	}
	if ep.Responses[0].Kind != "json" {
		t.Fatalf("kind: %q", ep.Responses[0].Kind)
	}
}

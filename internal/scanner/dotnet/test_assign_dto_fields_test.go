//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what TestAssignDTOFields 테스트
package dotnet

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestAssignDTOFields(t *testing.T) {
	body := &scanner.Endpoint{Request: &scanner.Request{Body: &scanner.Body{}}}
	assignDTOFields(dtoRequest{isBody: true}, body, []scanner.Field{{Name: "x"}})
	if len(body.Request.Body.Fields) != 1 {
		t.Fatalf("body: %+v", body.Request.Body)
	}
	resp := &scanner.Endpoint{Responses: []scanner.Response{{Status: "200"}}}
	assignDTOFields(dtoRequest{}, resp, []scanner.Field{{Name: "y"}})
	if len(resp.Responses[0].Fields) != 1 {
		t.Fatalf("resp: %+v", resp.Responses)
	}
}

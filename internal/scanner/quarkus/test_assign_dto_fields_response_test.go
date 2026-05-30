//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestAssignDTOFields_Response 테스트
package quarkus

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestAssignDTOFields_Response(t *testing.T) {
	ep := &scanner.Endpoint{Responses: []scanner.Response{{Status: "200"}}}
	fields := []scanner.Field{{Name: "name", Type: "string"}}
	assignDTOFields(dtoRequest{}, ep, fields)
	if len(ep.Responses[0].Fields) != 1 {
		t.Fatalf("got %+v", ep.Responses)
	}
}

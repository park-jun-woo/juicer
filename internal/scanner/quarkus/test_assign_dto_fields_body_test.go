//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestAssignDTOFields_Body 테스트
package quarkus

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestAssignDTOFields_Body(t *testing.T) {
	ep := &scanner.Endpoint{Request: &scanner.Request{Body: &scanner.Body{}}}
	fields := []scanner.Field{{Name: "name", Type: "string"}}
	assignDTOFields(dtoRequest{isBody: true}, ep, fields)
	if len(ep.Request.Body.Fields) != 1 {
		t.Fatalf("got %+v", ep.Request.Body)
	}
}

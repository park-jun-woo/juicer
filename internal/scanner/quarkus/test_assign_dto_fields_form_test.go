//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestAssignDTOFields_Form 테스트
package quarkus

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestAssignDTOFields_Form(t *testing.T) {
	ep := &scanner.Endpoint{}
	fields := []scanner.Field{{Name: "name", Type: "string"}}
	assignDTOFields(dtoRequest{isForm: true}, ep, fields)
	if ep.Request == nil || len(ep.Request.FormFields) != 1 {
		t.Fatalf("got %+v", ep.Request)
	}
}

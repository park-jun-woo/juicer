//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestApplyPrimitivePathType_Round5 테스트
package actix

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestApplyPrimitivePathType_Round5(t *testing.T) {
	ep := &scanner.Endpoint{}
	applyPathParams(ep, "/users/{id}")
	applyPrimitivePathType(ep, "integer")
	if ep.Request.PathParams[0].Type != "integer" {
		t.Fatalf("type: %q", ep.Request.PathParams[0].Type)
	}
}

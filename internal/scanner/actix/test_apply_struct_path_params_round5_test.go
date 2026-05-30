//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestApplyStructPathParams_Round5 테스트
package actix

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestApplyStructPathParams_Round5(t *testing.T) {
	ep := &scanner.Endpoint{}
	applyStructPathParams(ep, []scanner.Field{{JSON: "id", Type: "integer"}, {JSON: "slug", Type: "string"}})
	if len(ep.Request.PathParams) != 2 {
		t.Fatalf("params: %+v", ep.Request.PathParams)
	}
}

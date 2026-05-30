//ff:func feature=scan type=test control=sequence topic=zod
//ff:what TestApplyParamValidator 테스트
package zod

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestApplyParamValidator(t *testing.T) {
	req := &scanner.Request{PathParams: []scanner.Param{{Name: "id", Type: "string"}}}

	ApplyParamValidator(req, []scanner.Field{{Name: "id", Type: "integer"}})
	if req.PathParams[0].Type != "integer" {
		t.Fatalf("update: %+v", req.PathParams)
	}

	changed := ApplyParamValidator(req, []scanner.Field{{Name: "extra", Type: "string"}})
	if !changed || len(req.PathParams) != 2 {
		t.Fatalf("append: %+v", req.PathParams)
	}
}

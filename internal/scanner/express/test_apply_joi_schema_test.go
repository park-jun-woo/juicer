//ff:func feature=scan type=test control=sequence topic=express
//ff:what applyJoiSchema body/query/params 반영 및 changed 집계 테스트
package express

import (
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
	"github.com/park-jun-woo/codistill/internal/scanner/joi"
)

func TestApplyJoiSchema(t *testing.T) {
	req := &scanner.Request{}
	rs := joi.RequestSchema{
		Body:   []scanner.Field{{Name: "name", Type: "string"}},
		Query:  []scanner.Field{{Name: "page", Type: "integer"}},
		Params: []scanner.Field{{Name: "id", Type: "string"}},
	}
	if !applyJoiSchema(req, rs) {
		t.Fatal("should report changed")
	}
	if req.Body == nil || req.Body.Method != "json" || len(req.Body.Fields) != 1 {
		t.Errorf("body: %+v", req.Body)
	}
	if len(req.Query) != 1 || len(req.PathParams) != 1 {
		t.Errorf("query/params: %+v %+v", req.Query, req.PathParams)
	}
	// empty schema -> no change
	if applyJoiSchema(&scanner.Request{}, joi.RequestSchema{}) {
		t.Error("empty schema should be unchanged")
	}
}

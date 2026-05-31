//ff:func feature=scan type=test topic=express control=sequence
//ff:what applyJoiQuery 쿼리 필드 추가 및 changed 반환 테스트
package express

import (
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestApplyJoiQuery(t *testing.T) {
	req := &scanner.Request{}
	if applyJoiQuery(req, nil) {
		t.Error("empty fields -> false")
	}
	fields := []scanner.Field{{Name: "page", Type: "integer"}}
	if !applyJoiQuery(req, fields) {
		t.Error("non-empty -> true")
	}
	if len(req.Query) != 1 || req.Query[0].Name != "page" || req.Query[0].Type != "integer" {
		t.Errorf("query: %+v", req.Query)
	}
}

//ff:func feature=scan type=test topic=express control=sequence
//ff:what applyJoiParams 중복 제외 path param 추가 및 changed 반환 테스트
package express

import (
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestApplyJoiParams(t *testing.T) {
	req := &scanner.Request{PathParams: []scanner.Param{{Name: "id"}}}
	fields := []scanner.Field{{Name: "id", Type: "string"}, {Name: "slug", Type: "string"}}
	if !applyJoiParams(req, fields) {
		t.Error("new param -> true")
	}
	if len(req.PathParams) != 2 {
		t.Errorf("dup id must be skipped: %+v", req.PathParams)
	}
	// no new fields
	if applyJoiParams(req, []scanner.Field{{Name: "id"}}) {
		t.Error("all existing -> false")
	}
}

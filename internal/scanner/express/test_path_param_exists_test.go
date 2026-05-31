//ff:func feature=scan type=test topic=express control=sequence
//ff:what pathParamExists 파라미터 이름 존재 여부 테스트
package express

import (
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestPathParamExists(t *testing.T) {
	params := []scanner.Param{{Name: "id"}, {Name: "slug"}}
	if !pathParamExists(params, "id") {
		t.Error("id should exist")
	}
	if pathParamExists(params, "missing") {
		t.Error("missing should not exist")
	}
	if pathParamExists(nil, "id") {
		t.Error("nil params -> false")
	}
}

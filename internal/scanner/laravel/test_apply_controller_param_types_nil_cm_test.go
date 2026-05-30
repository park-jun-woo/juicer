//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestApplyControllerParamTypes_NilCM 테스트
package laravel

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestApplyControllerParamTypes_NilCM(t *testing.T) {
	pp := []scanner.Param{{Name: "id", Type: "string"}}
	got := applyControllerParamTypes(pp, nil)
	if len(got) != 1 || got[0].Type != "string" {
		t.Fatalf("got %+v", got)
	}
}

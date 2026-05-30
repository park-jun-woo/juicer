//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestApplyControllerParamTypes_Typed 테스트
package laravel

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestApplyControllerParamTypes_Typed(t *testing.T) {
	cm := &controllerMethod{params: []methodParam{{name: "id", typeName: "int"}}}
	pp := []scanner.Param{{Name: "id", Type: "string"}, {Name: "slug", Type: "string"}}
	got := applyControllerParamTypes(pp, cm)
	if got[0].Type != "integer" {
		t.Fatalf("id type: %+v", got[0])
	}
	if got[1].Type != "string" {
		t.Fatalf("slug unchanged: %+v", got[1])
	}
}

//ff:func feature=scan type=test control=sequence topic=zod
//ff:what TestFindPathParamIndex 테스트
package zod

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestFindPathParamIndex(t *testing.T) {
	params := []scanner.Param{{Name: "id"}, {Name: "slug"}}
	if findPathParamIndex(params, "slug") != 1 {
		t.Fatal("found")
	}
	if findPathParamIndex(params, "missing") != -1 {
		t.Fatal("not found")
	}
}

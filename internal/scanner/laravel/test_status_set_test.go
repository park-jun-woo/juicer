//ff:func feature=scan type=test topic=laravel control=sequence
//ff:what statusSet 응답 status 집합 생성 테스트
package laravel

import (
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestStatusSet(t *testing.T) {
	set := statusSet([]scanner.Response{{Status: "200"}, {Status: "404"}, {Status: "200"}})
	if !set["200"] || !set["404"] {
		t.Errorf("set: %v", set)
	}
	if len(set) != 2 {
		t.Errorf("len = %d, want 2", len(set))
	}
}

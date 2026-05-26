//ff:func feature=scan type=extract control=sequence
//ff:what TestExtractRoutes_Empty 테스트
package scanner

import (
	"testing"
)

func TestExtractRoutes_Empty(t *testing.T) {
	eps, _ := extractRoutes(nil, ".")
	if len(eps) != 0 {
		t.Errorf("expected 0, got %d", len(eps))
	}
}

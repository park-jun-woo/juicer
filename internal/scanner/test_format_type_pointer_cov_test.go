//ff:func feature=scan type=test control=sequence
//ff:what TestFormatType_PointerCov 테스트
package scanner

import (
	"go/types"
	"testing"
)

func TestFormatType_PointerCov(t *testing.T) {
	got := formatType(types.NewPointer(types.Typ[types.Int]))
	if got != "*int" {
		t.Fatalf("expected *int, got %s", got)
	}
}

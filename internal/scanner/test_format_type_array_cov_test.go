//ff:func feature=scan type=test control=sequence
//ff:what TestFormatType_ArrayCov 테스트
package scanner

import (
	"go/types"
	"testing"
)

func TestFormatType_ArrayCov(t *testing.T) {
	got := formatType(types.NewArray(types.Typ[types.Int], 10))
	if got != "[]int" {
		t.Fatalf("expected []int, got %s", got)
	}
}

//ff:func feature=scan type=extract control=sequence
//ff:what TestFormatType_Array 테스트
package gogin

import (
	"go/types"
	"testing"
)

func TestFormatType_Array(t *testing.T) {
	got := formatType(types.NewArray(types.Typ[types.Int], 5))
	if got != "[]int" {
		t.Fatalf("expected []int, got %s", got)
	}
}

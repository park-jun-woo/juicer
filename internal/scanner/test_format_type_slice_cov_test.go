//ff:func feature=scan type=test control=sequence
//ff:what TestFormatType_SliceCov 테스트
package scanner

import (
	"go/types"
	"testing"
)

func TestFormatType_SliceCov(t *testing.T) {
	got := formatType(types.NewSlice(types.Typ[types.String]))
	if got != "[]string" {
		t.Fatalf("expected []string, got %s", got)
	}
}

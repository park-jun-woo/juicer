//ff:func feature=scan type=extract control=sequence
//ff:what TestFormatType_Slice 테스트
package scanner

import (
	"go/types"
	"testing"
)

func TestFormatType_Slice(t *testing.T) {
	got := formatType(types.NewSlice(types.Typ[types.String]))
	if got != "[]string" {
		t.Fatalf("expected []string, got %s", got)
	}
}

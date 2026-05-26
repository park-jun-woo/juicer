//ff:func feature=scan type=extract control=sequence
//ff:what TestFormatType_Pointer 테스트
package gogin

import (
	"go/types"
	"testing"
)

func TestFormatType_Pointer(t *testing.T) {
	got := formatType(types.NewPointer(types.Typ[types.Int]))
	if got != "*int" {
		t.Fatalf("expected *int, got %s", got)
	}
}

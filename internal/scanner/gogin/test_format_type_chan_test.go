//ff:func feature=scan type=extract control=sequence
//ff:what TestFormatType_Chan 테스트
package gogin

import (
	"go/types"
	"testing"
)

func TestFormatType_Chan(t *testing.T) {
	ch := types.NewChan(types.SendRecv, types.Typ[types.Int])
	got := formatType(ch)
	// Chan is not a handled case, uses t.String()
	if got == "" {
		t.Error("expected non-empty string")
	}
}

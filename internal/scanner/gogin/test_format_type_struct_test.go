//ff:func feature=scan type=extract control=sequence
//ff:what TestFormatType_Struct 테스트
package gogin

import (
	"go/types"
	"testing"
)

func TestFormatType_Struct(t *testing.T) {
	st := types.NewStruct(nil, nil)
	got := formatType(st)
	if got != "object" {
		t.Errorf("expected 'object', got %q", got)
	}
}

//ff:func feature=scan type=extract control=sequence
//ff:what TestFormatType_NamedNilPkg 테스트
package gogin

import (
	"go/types"
	"testing"
)

func TestFormatType_NamedNilPkg(t *testing.T) {
	named := types.NewNamed(
		types.NewTypeName(0, nil, "error", nil),
		types.NewInterfaceType(nil, nil),
		nil,
	)
	got := formatType(named)
	if got != "error" {
		t.Errorf("expected 'error', got %q", got)
	}
}

//ff:func feature=scan type=extract control=sequence
//ff:what TestFormatType_Interface 테스트
package gogin

import (
	"go/types"
	"testing"
)

func TestFormatType_Interface(t *testing.T) {
	got := formatType(types.NewInterfaceType(nil, nil))
	if got != "any" {
		t.Fatalf("expected any, got %s", got)
	}
}

//ff:func feature=scan type=test control=sequence
//ff:what TestFormatType_InterfaceCov 테스트
package gogin

import (
	"go/types"
	"testing"
)

func TestFormatType_InterfaceCov(t *testing.T) {
	got := formatType(types.NewInterfaceType(nil, nil))
	if got != "any" {
		t.Fatalf("expected any, got %s", got)
	}
}

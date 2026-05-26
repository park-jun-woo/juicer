//ff:func feature=scan type=test control=sequence
//ff:what TestFormatType_Basic 테스트
package scanner

import (
	"go/types"
	"testing"
)

func TestFormatType_Basic(t *testing.T) {
	got := formatType(types.Typ[types.String])
	if got != "string" {
		t.Fatalf("expected string, got %s", got)
	}
}


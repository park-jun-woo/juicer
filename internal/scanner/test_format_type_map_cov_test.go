//ff:func feature=scan type=test control=sequence
//ff:what TestFormatType_MapCov 테스트
package scanner

import (
	"go/types"
	"testing"
)

func TestFormatType_MapCov(t *testing.T) {
	got := formatType(types.NewMap(types.Typ[types.String], types.Typ[types.Int]))
	if got != "map[string]int" {
		t.Fatalf("expected map[string]int, got %s", got)
	}
}

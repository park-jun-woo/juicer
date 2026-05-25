//ff:func feature=scan type=extract control=sequence
//ff:what TestExtractFields_Empty 테스트
package scanner

import (
	"go/types"
	"testing"
)

func TestExtractFields_Empty(t *testing.T) {
	st := types.NewStruct(nil, nil)
	result := extractFields(st, make(map[string]bool))
	if len(result) != 0 {
		t.Fatalf("expected 0 fields, got %d", len(result))
	}
}

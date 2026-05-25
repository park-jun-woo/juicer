//ff:func feature=scan type=extract control=sequence
//ff:what TestFormatType_StructType 테스트
package scanner

import (
	"go/types"
	"testing"
)

func TestFormatType_StructType(t *testing.T) {
	got := formatType(types.NewStruct(nil, nil))
	if got != "object" {
		t.Fatalf("expected object, got %s", got)
	}
}

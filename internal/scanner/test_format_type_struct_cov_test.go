//ff:func feature=scan type=test control=sequence
//ff:what TestFormatType_StructCov 테스트
package scanner

import (
	"go/types"
	"testing"
)

func TestFormatType_StructCov(t *testing.T) {
	got := formatType(types.NewStruct(nil, nil))
	if got != "object" {
		t.Fatalf("expected object, got %s", got)
	}
}

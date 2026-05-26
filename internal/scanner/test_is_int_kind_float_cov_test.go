//ff:func feature=scan type=test control=sequence
//ff:what TestIsIntKind_FloatCov 테스트
package scanner

import (
	"go/types"
	"testing"
)

func TestIsIntKind_FloatCov(t *testing.T) {
	if isIntKind(types.Float64) {
		t.Fatal("expected false for float64")
	}
}

//ff:func feature=scan type=test control=sequence topic=echo
//ff:what TestIsIntKind 테스트
package echo

import (
	"go/types"
	"testing"
)

func TestIsIntKind(t *testing.T) {
	if !isIntKind(types.Int) || !isIntKind(types.Uint64) {
		t.Fatal("int kinds")
	}
	if isIntKind(types.String) || isIntKind(types.Float64) {
		t.Fatal("non-int kinds")
	}
}

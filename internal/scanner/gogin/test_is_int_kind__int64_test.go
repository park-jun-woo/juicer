//ff:func feature=scan type=extract control=sequence
//ff:what TestIsIntKind_Int64 테스트
package gogin

import (
	"go/types"
	"testing"
)

func TestIsIntKind_Int64(t *testing.T) {
	if !isIntKind(types.Int64) {
		t.Fatal("expected true")
	}
}

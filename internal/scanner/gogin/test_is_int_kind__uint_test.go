//ff:func feature=scan type=extract control=sequence
//ff:what TestIsIntKind_Uint 테스트
package gogin

import (
	"go/types"
	"testing"
)

func TestIsIntKind_Uint(t *testing.T) {
	if !isIntKind(types.Uint) {
		t.Fatal("expected true")
	}
}

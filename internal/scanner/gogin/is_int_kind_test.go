//ff:func feature=scan type=test control=sequence
//ff:what TestIsIntKind_Int 테스트
package gogin

import (
	"go/types"
	"testing"
)

func TestIsIntKind_Int(t *testing.T) {
	if !isIntKind(types.Int) {
		t.Fatal("expected true")
	}
}


//ff:func feature=scan type=extract control=sequence
//ff:what TestIsIntKind_String 테스트
package scanner

import (
	"go/types"
	"testing"
)

func TestIsIntKind_String(t *testing.T) {
	if isIntKind(types.String) {
		t.Fatal("expected false")
	}
}

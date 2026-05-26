//ff:func feature=scan type=test control=sequence
//ff:what TestFindFileForPos_NilPkgs 테스트
package gogin

import (
	"go/token"
	"testing"
)

func TestFindFileForPos_NilPkgs(t *testing.T) {
	if got := findFileForPos(token.NoPos, nil); got != nil {
		t.Fatal("expected nil for nil pkgs")
	}
}

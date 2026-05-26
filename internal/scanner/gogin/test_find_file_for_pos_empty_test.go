//ff:func feature=scan type=test control=sequence
//ff:what TestFindFileForPos_EmptyPkgs 테스트
package gogin

import (
	"go/token"
	"testing"

	"golang.org/x/tools/go/packages"
)

func TestFindFileForPos_EmptyPkgs(t *testing.T) {
	if got := findFileForPos(token.NoPos, []*packages.Package{}); got != nil {
		t.Fatal("expected nil for empty pkgs")
	}
}

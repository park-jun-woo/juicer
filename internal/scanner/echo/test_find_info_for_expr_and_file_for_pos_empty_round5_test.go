//ff:func feature=scan type=test control=sequence topic=echo
//ff:what TestFindInfoForExpr_And_FileForPos_Empty_Round5 테스트
package echo

import (
	"go/token"
	"testing"
)

func TestFindInfoForExpr_And_FileForPos_Empty_Round5(t *testing.T) {
	expr := parseExpr(t, "x")
	if info := findInfoForExpr(expr, nil); info != nil {
		t.Fatal("expected nil info for empty pkgs")
	}
	if f := findFileForPos(token.Pos(1), nil); f != nil {
		t.Fatal("expected nil file for empty pkgs")
	}
}

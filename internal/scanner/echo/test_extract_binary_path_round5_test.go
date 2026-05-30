//ff:func feature=scan type=test control=sequence topic=echo
//ff:what TestExtractBinaryPath_Round5 테스트
package echo

import (
	"go/ast"
	"testing"
)

func TestExtractBinaryPath_Round5(t *testing.T) {

	be := parseExpr(t, `"/a" + "/b"`).(*ast.BinaryExpr)
	got, ok := extractBinaryPath(nil, be)
	if !ok || got != "/a/b" {
		t.Fatalf("got %q %v", got, ok)
	}
}

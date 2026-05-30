//ff:func feature=scan type=test control=sequence
//ff:what TestResolveStatusCode_Unknown 테스트
package fiber

import (
	"go/ast"
	"testing"
)

func TestResolveStatusCode_Unknown(t *testing.T) {

	if got := resolveStatusCode(&ast.Ident{Name: "dynamic"}, newEmptyInfoFull()); got != "(unknown)" {
		t.Fatalf("unknown: got %q", got)
	}
}

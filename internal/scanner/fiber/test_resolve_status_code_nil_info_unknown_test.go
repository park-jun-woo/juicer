//ff:func feature=scan type=test control=sequence
//ff:what TestResolveStatusCode_NilInfoUnknown 테스트
package fiber

import (
	"go/ast"
	"testing"
)

func TestResolveStatusCode_NilInfoUnknown(t *testing.T) {
	if got := resolveStatusCode(&ast.Ident{Name: "x"}, nil); got != "(unknown)" {
		t.Fatalf("nil info: got %q", got)
	}
}

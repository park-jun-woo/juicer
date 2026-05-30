//ff:func feature=scan type=test control=sequence
//ff:what TestResolveResponseType_NilInfo 테스트
package fiber

import (
	"go/ast"
	"testing"
)

func TestResolveResponseType_NilInfo(t *testing.T) {
	tn, f, c := resolveResponseType(&ast.Ident{Name: "x"}, nil)
	if tn != "" || f != nil || c != "" {
		t.Fatalf("nil info: %q %v %q", tn, f, c)
	}
}

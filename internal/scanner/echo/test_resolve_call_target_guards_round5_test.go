//ff:func feature=scan type=test control=sequence topic=echo
//ff:what TestResolveCallTarget_Guards_Round5 테스트
package echo

import (
	"go/ast"
	"testing"
)

func TestResolveCallTarget_Guards_Round5(t *testing.T) {
	file, info := checkSrc(t, `package m
func Target() int { return 0 }
var _ = Target()
`)
	var call *ast.CallExpr
	ast.Inspect(file, func(n ast.Node) bool {
		if c, ok := n.(*ast.CallExpr); ok {
			call = c
		}
		return true
	})
	if call == nil {
		t.Fatal("no call")
	}
	pos := resolveCallTarget(call, info)
	if !pos.IsValid() {
		t.Fatal("expected valid target pos")
	}
}

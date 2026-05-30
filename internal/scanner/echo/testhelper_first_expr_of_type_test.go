//ff:func feature=scan type=test control=sequence topic=echo
//ff:what firstExprOfType 테스트 헬퍼
package echo

import (
	"go/ast"
	"testing"
)

func firstExprOfType[T ast.Node](t *testing.T, file *ast.File, pred func(T) bool) T {
	t.Helper()
	var found T
	var ok bool
	ast.Inspect(file, func(n ast.Node) bool {
		if ok {
			return false
		}
		if tn, is := n.(T); is && pred(tn) {
			found = tn
			ok = true
			return false
		}
		return true
	})
	if !ok {
		t.Fatal("node not found")
	}
	return found
}

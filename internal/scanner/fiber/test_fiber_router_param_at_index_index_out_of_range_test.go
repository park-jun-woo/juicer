//ff:func feature=scan type=test control=sequence
//ff:what TestFiberRouterParamAtIndex_IndexOutOfRange 테스트
package fiber

import "testing"

func TestFiberRouterParamAtIndex_IndexOutOfRange(t *testing.T) {
	fn := funcDeclFrom(t, "package m\nfunc Setup(app int) {}\n")
	if got := fiberRouterParamAtIndex(fn, newEmptyInfo(), 5); got != "" {
		t.Fatalf("out of range: got %q", got)
	}
}

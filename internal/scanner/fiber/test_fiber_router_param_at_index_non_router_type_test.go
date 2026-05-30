//ff:func feature=scan type=test control=sequence
//ff:what TestFiberRouterParamAtIndex_NonRouterType 테스트
package fiber

import "testing"

func TestFiberRouterParamAtIndex_NonRouterType(t *testing.T) {

	fn := funcDeclFrom(t, "package m\nfunc Setup(app int) {}\n")
	if got := fiberRouterParamAtIndex(fn, newEmptyInfo(), 0); got != "" {
		t.Fatalf("non-router: got %q", got)
	}
}

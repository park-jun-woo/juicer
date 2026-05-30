//ff:func feature=scan type=test control=sequence
//ff:what TestFiberRouterParamAtIndex_SecondParam 테스트
package fiber

import "testing"

func TestFiberRouterParamAtIndex_SecondParam(t *testing.T) {

	fn := funcDeclFrom(t, "package m\nfunc Setup(a int, b string) {}\n")
	if got := fiberRouterParamAtIndex(fn, newEmptyInfo(), 1); got != "" {
		t.Fatalf("second param: got %q", got)
	}
}

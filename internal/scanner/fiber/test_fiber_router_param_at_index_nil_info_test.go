//ff:func feature=scan type=test control=sequence
//ff:what TestFiberRouterParamAtIndex_NilInfo 테스트
package fiber

import "testing"

func TestFiberRouterParamAtIndex_NilInfo(t *testing.T) {
	fn := funcDeclFrom(t, "package m\nfunc Setup(app int) {}\n")
	if got := fiberRouterParamAtIndex(fn, nil, 0); got != "" {
		t.Fatalf("nil info: got %q", got)
	}
}

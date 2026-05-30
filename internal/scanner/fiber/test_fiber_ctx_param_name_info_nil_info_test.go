//ff:func feature=scan type=test control=sequence
//ff:what TestFiberCtxParamNameInfo_NilInfo 테스트
package fiber

import "testing"

func TestFiberCtxParamNameInfo_NilInfo(t *testing.T) {

	if got := fiberCtxParamNameInfo(ctxFuncType("c"), nil); got != "c" {
		t.Fatalf("nil info: got %q, want c", got)
	}
}

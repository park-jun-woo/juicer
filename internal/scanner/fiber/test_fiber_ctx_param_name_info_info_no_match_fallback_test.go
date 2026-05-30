//ff:func feature=scan type=test control=sequence
//ff:what TestFiberCtxParamNameInfo_InfoNoMatchFallback 테스트
package fiber

import "testing"

func TestFiberCtxParamNameInfo_InfoNoMatchFallback(t *testing.T) {

	if got := fiberCtxParamNameInfo(ctxFuncType("ctx"), newEmptyInfo()); got != "ctx" {
		t.Fatalf("fallback: got %q, want ctx", got)
	}
}

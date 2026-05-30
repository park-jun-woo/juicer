//ff:func feature=scan type=test control=sequence topic=echo
//ff:what TestRescanCalleeWithPrefix_DepthGuard_Round5 테스트
package echo

import "testing"

func TestRescanCalleeWithPrefix_DepthGuard_Round5(t *testing.T) {
	ctx := emptyGroupCtx()
	call := callExprFrom(t, `target(g)`)

	rescanCalleeWithPrefixDepth(call, 0, "/p", &routerInfo{}, ctx, maxRescanDepth)
}

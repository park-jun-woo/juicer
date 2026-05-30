//ff:func feature=scan type=test control=sequence topic=echo
//ff:what TestCheckOneDepthCall_Guards_Round5 테스트
package echo

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestCheckOneDepthCall_Guards_Round5(t *testing.T) {
	ep := &scanner.Endpoint{}

	checkOneDepthCall(ep, callExprFrom(t, `helper(c)`), "c", nil, buildFuncIndex(nil))
}

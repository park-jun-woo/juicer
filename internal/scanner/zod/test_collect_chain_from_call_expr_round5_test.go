//ff:func feature=scan type=test control=sequence topic=zod
//ff:what TestCollectChainFromCallExpr_Round5 테스트
package zod

import "testing"

func TestCollectChainFromCallExpr_Round5(t *testing.T) {
	root, src := parseTS(t, "const s = z.string().min(2);")
	call := findFirstCall(t, root)
	var methods []ChainMethod
	collectChainFromCallExpr(call, src, &methods)
	if len(methods) == 0 {
		t.Fatalf("expected methods, got %+v", methods)
	}
}

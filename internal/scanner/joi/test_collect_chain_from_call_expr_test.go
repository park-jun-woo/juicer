//ff:func feature=scan type=test topic=joi control=sequence
//ff:what collectChainFromCallExpr 메서드 호출에서 ChainMethod 추가 직접 테스트
package joi

import "testing"

func TestCollectChainFromCallExpr(t *testing.T) {
	root, src := parseJoiTS(t, `Joi.string().email()`)
	// the .email() call_expression
	call := firstOfType(root, "call_expression")
	var methods []ChainMethod
	collectChainFromCallExpr(call, src, &methods)
	if len(methods) == 0 {
		t.Fatalf("expected at least one method")
	}
	// non-member function: f()
	root2, src2 := parseJoiTS(t, `f()`)
	call2 := firstOfType(root2, "call_expression")
	var m2 []ChainMethod
	collectChainFromCallExpr(call2, src2, &m2)
	if len(m2) != 0 {
		t.Errorf("plain call should add nothing: %v", m2)
	}
}

//ff:func feature=scan type=test topic=joi control=sequence
//ff:what collectChainMethodsRecursive call/member 분기 및 nil 가드 직접 테스트
package joi

import "testing"

func TestCollectChainMethodsRecursive(t *testing.T) {
	// nil guard
	var methods []ChainMethod
	collectChainMethodsRecursive(nil, nil, &methods)
	if len(methods) != 0 {
		t.Errorf("nil node should add nothing: %v", methods)
	}

	// call_expression branch via Joi.string()
	root, src := parseJoiTS(t, `Joi.string()`)
	call := firstOfType(root, "call_expression")
	methods = nil
	collectChainMethodsRecursive(call, src, &methods)
	if len(methods) == 0 {
		t.Fatalf("expected methods from call expr")
	}
}

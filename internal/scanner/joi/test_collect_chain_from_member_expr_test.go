//ff:func feature=scan type=test topic=joi control=sequence
//ff:what collectChainFromMemberExpr Joi 베이스 메서드 추가/재귀 직접 테스트
package joi

import "testing"

func TestCollectChainFromMemberExpr(t *testing.T) {
	// Joi.string : member_expression with object "Joi"
	root, src := parseJoiTS(t, `Joi.string`)
	mem := firstOfType(root, "member_expression")
	if mem == nil {
		t.Fatal("no member_expression")
	}
	var methods []ChainMethod
	collectChainFromMemberExpr(mem, src, &methods)
	if len(methods) != 1 || methods[0].Name != "string" {
		t.Fatalf("Joi.string: %v", methods)
	}
}

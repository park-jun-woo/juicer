//ff:func feature=scan type=test control=sequence topic=joi
//ff:what ParseChain 메서드 체인 → Field(기본 string) 변환 테스트
package joi

import "testing"

func TestParseChain(t *testing.T) {
	root, src := parseJoiTS(t, `Joi.string().email().required()`)
	call := firstOfType(root, "call_expression")
	f := ParseChain(call, src)
	if f.Type != "string" {
		t.Errorf("type: %q", f.Type)
	}
	if f.Validate == "" {
		t.Errorf("validate should include email/required: %q", f.Validate)
	}

	// untyped chain defaults to string
	root2, src2 := parseJoiTS(t, `Joi.required()`)
	call2 := firstOfType(root2, "call_expression")
	if g := ParseChain(call2, src2); g.Type != "string" {
		t.Errorf("default type: %q", g.Type)
	}
}

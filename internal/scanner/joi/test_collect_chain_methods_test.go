//ff:func feature=scan type=test topic=joi control=iteration dimension=1
//ff:what CollectChainMethods 메서드 체인 inner→outer 수집 테스트
package joi

import "testing"

func TestCollectChainMethods(t *testing.T) {
	// Joi.string().email().required()
	root, src := parseJoiTS(t, `Joi.string().email().required()`)
	// outermost call_expression is the .required() call
	call := firstOfType(root, "call_expression")
	methods := CollectChainMethods(call, src)
	var names []string
	for _, m := range methods {
		names = append(names, m.Name)
	}
	// inner -> outer: string, email, required
	if len(names) < 3 || names[0] != "string" {
		t.Fatalf("got %v", names)
	}
	last := names[len(names)-1]
	if last != "required" {
		t.Errorf("last should be required, got %v", names)
	}
}

//ff:func feature=scan type=test topic=joi control=sequence
//ff:what FindKeysObject Joi.object().keys({...}) 필드 object 노드 탐색 테스트
package joi

import "testing"

func TestFindKeysObject(t *testing.T) {
	root, src := parseJoiTS(t, `Joi.object().keys({ name: Joi.string() })`)
	call := firstOfType(root, "call_expression")
	obj := FindKeysObject(call, src)
	if obj == nil || obj.Type() != "object" {
		t.Fatalf("keys object not found: %v", obj)
	}
	// no schema method -> nil
	root2, src2 := parseJoiTS(t, `foo.bar({ x: 1 })`)
	call2 := firstOfType(root2, "call_expression")
	if FindKeysObject(call2, src2) != nil {
		t.Error("non-joi method should be nil")
	}
}

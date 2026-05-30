//ff:func feature=scan type=test control=sequence topic=zod
//ff:what TestFindAllByTypeAndChild 테스트
package zod

import "testing"

func TestFindAllByTypeAndChild(t *testing.T) {
	root, _ := parseTS(t, `a(); b();`)
	if len(findAllByType(root, "call_expression")) != 2 {
		t.Fatal("findAllByType")
	}
	calls := findAllByType(root, "call_expression")
	if findChildByType(calls[0], "arguments") == nil {
		t.Fatal("findChildByType")
	}
	if findChildByType(calls[0], "object") != nil {
		t.Fatal("nil expected")
	}
}

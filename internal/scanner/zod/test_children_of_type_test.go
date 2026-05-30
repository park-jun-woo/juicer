//ff:func feature=scan type=test control=sequence topic=zod
//ff:what TestChildrenOfType 테스트
package zod

import "testing"

func TestChildrenOfType(t *testing.T) {
	root, _ := parseTS(t, `const o = { a: 1, b: 2 };`)
	objs := findAllByType(root, "object")
	if len(childrenOfType(objs[0], "pair")) != 2 {
		t.Fatal("children")
	}
}

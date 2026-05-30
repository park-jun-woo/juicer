//ff:func feature=scan type=test control=sequence topic=supafunc
//ff:what TestChildrenOfType 테스트
package supafunc

import "testing"

func TestChildrenOfType(t *testing.T) {
	fi := mustParse(t, []byte(`const o = { a: 1, b: 2 };`))
	objs := findAllByType(fi.Root, "object")
	if len(childrenOfType(objs[0], "pair")) != 2 {
		t.Fatal("children")
	}
}

//ff:func feature=scan type=test control=sequence topic=zod
//ff:what TestCollectArgNodes 테스트
package zod

import "testing"

func TestCollectArgNodes(t *testing.T) {
	root, _ := parseTS(t, `f("a", b, 1);`)
	args := findAllByType(root, "arguments")[0]
	nodes := collectArgNodes(args)
	if len(nodes) != 3 {
		t.Fatalf("got %d", len(nodes))
	}
}

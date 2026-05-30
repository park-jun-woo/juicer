//ff:func feature=scan type=test control=sequence topic=zod
//ff:what TestExtractArgValues 테스트
package zod

import "testing"

func TestExtractArgValues(t *testing.T) {
	root, src := parseTS(t, `f("str", 42, ['a','b']);`)
	args := findAllByType(root, "arguments")[0]
	nodes := collectArgNodes(args)
	if got := extractArgValues(nodes[0], src); len(got) != 1 || got[0] != "str" {
		t.Fatalf("string: %v", got)
	}
	if got := extractArgValues(nodes[1], src); len(got) != 1 || got[0] != "42" {
		t.Fatalf("number: %v", got)
	}
	if got := extractArgValues(nodes[2], src); len(got) != 2 {
		t.Fatalf("array: %v", got)
	}
}

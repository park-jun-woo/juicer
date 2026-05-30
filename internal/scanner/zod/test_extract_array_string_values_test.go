//ff:func feature=scan type=test control=sequence topic=zod
//ff:what TestExtractArrayStringValues 테스트
package zod

import "testing"

func TestExtractArrayStringValues(t *testing.T) {
	root, src := parseTS(t, `const x = ['a', 'b', 'c'];`)
	arr := findAllByType(root, "array")[0]
	vals := extractArrayStringValues(arr, src)
	if len(vals) != 3 || vals[0] != "a" {
		t.Fatalf("got %v", vals)
	}
}

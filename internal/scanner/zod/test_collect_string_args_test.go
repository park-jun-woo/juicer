//ff:func feature=scan type=test control=sequence topic=zod
//ff:what TestCollectStringArgs 테스트
package zod

import "testing"

func TestCollectStringArgs(t *testing.T) {
	root, src := parseTS(t, `z.enum(['a', 'b']);`)
	args := findAllByType(root, "arguments")[0]
	got := collectStringArgs(args, src)
	if len(got) != 2 {
		t.Fatalf("got %v", got)
	}
}

//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestReadObjectStringProp_NotString 테스트
package hono

import "testing"

func TestReadObjectStringProp_NotString(t *testing.T) {

	fi := mustParse(t, []byte(`const o = { count: 3 };`+"\n"))
	obj := findAllByType(fi.Root, "object")[0]
	if got := readObjectStringProp(obj, "count", fi.Src); got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}

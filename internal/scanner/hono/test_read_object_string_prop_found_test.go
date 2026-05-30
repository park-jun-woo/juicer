//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestReadObjectStringProp_Found 테스트
package hono

import "testing"

func TestReadObjectStringProp_Found(t *testing.T) {
	fi := mustParse(t, []byte(`const o = { method: "post", count: 3 };`+"\n"))
	obj := findAllByType(fi.Root, "object")[0]
	if got := readObjectStringProp(obj, "method", fi.Src); got != "post" {
		t.Fatalf("got %q", got)
	}
}

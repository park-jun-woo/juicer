//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestReadObjectStringProp_Missing 테스트
package hono

import "testing"

func TestReadObjectStringProp_Missing(t *testing.T) {
	fi := mustParse(t, []byte(`const o = { a: "x" };`+"\n"))
	obj := findAllByType(fi.Root, "object")[0]
	if got := readObjectStringProp(obj, "missing", fi.Src); got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}

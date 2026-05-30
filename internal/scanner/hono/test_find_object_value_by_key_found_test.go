//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestFindObjectValueByKey_Found 테스트
package hono

import "testing"

func TestFindObjectValueByKey_Found(t *testing.T) {
	fi := mustParse(t, []byte(`const o = { method: "get", path: "/x" };`+"\n"))
	obj := findAllByType(fi.Root, "object")[0]
	v := findObjectValueByKey(obj, "path", fi.Src)
	if v == nil || nodeText(v, fi.Src) != `"/x"` {
		t.Fatalf("got %v", v)
	}
}

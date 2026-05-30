//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestFindObjectValueByKey_NotFound 테스트
package hono

import "testing"

func TestFindObjectValueByKey_NotFound(t *testing.T) {
	fi := mustParse(t, []byte(`const o = { a: 1 };`+"\n"))
	obj := findAllByType(fi.Root, "object")[0]
	if findObjectValueByKey(obj, "missing", fi.Src) != nil {
		t.Fatal("expected nil")
	}
}

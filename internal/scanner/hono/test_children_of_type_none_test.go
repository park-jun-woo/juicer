//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestChildrenOfType_None 테스트
package hono

import "testing"

func TestChildrenOfType_None(t *testing.T) {
	fi := mustParse(t, []byte(`const o = {};`+"\n"))
	obj := findAllByType(fi.Root, "object")[0]
	if got := childrenOfType(obj, "pair"); len(got) != 0 {
		t.Fatalf("expected 0, got %d", len(got))
	}
}

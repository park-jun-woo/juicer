//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestFindAllByType_None 테스트
package hono

import "testing"

func TestFindAllByType_None(t *testing.T) {
	fi := mustParse(t, []byte(`const x = 1;`+"\n"))
	if got := findAllByType(fi.Root, "call_expression"); len(got) != 0 {
		t.Fatalf("expected 0, got %d", len(got))
	}
}

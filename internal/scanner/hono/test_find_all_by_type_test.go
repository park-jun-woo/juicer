//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestFindAllByType 테스트
package hono

import "testing"

func TestFindAllByType(t *testing.T) {
	fi := mustParse(t, []byte(`a(); b(); c();`+"\n"))
	calls := findAllByType(fi.Root, "call_expression")
	if len(calls) != 3 {
		t.Fatalf("expected 3 calls, got %d", len(calls))
	}
}

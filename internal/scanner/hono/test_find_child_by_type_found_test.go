//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestFindChildByType_Found 테스트
package hono

import "testing"

func TestFindChildByType_Found(t *testing.T) {
	fi := mustParse(t, []byte(`f();`+"\n"))
	call := findAllByType(fi.Root, "call_expression")[0]
	if findChildByType(call, "arguments") == nil {
		t.Fatal("expected arguments child")
	}
}

//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestFindChildByType_NotFound 테스트
package hono

import "testing"

func TestFindChildByType_NotFound(t *testing.T) {
	fi := mustParse(t, []byte(`f();`+"\n"))
	call := findAllByType(fi.Root, "call_expression")[0]
	if findChildByType(call, "object") != nil {
		t.Fatal("expected nil for missing type")
	}
}

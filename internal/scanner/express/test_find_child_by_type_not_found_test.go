//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestFindChildByType_NotFound 테스트
package express

import "testing"

func TestFindChildByType_NotFound(t *testing.T) {
	fi := mustParse(t, []byte(`f(a);`))
	call := firstCallExpr(t, fi)
	if got := findChildByType(call, "statement_block"); got != nil {
		t.Fatalf("expected nil, got %v", got.Type())
	}
}

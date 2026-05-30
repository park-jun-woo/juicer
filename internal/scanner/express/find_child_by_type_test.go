//ff:func feature=scan type=test control=sequence topic=express
//ff:what findChildByType: 일치 자식 반환 / 없으면 nil
package express

import "testing"

func TestFindChildByType_Found(t *testing.T) {
	fi := mustParse(t, []byte(`f(a);`))
	call := firstCallExpr(t, fi)
	if got := findChildByType(call, "arguments"); got == nil {
		t.Fatal("expected arguments child")
	}
}

func TestFindChildByType_NotFound(t *testing.T) {
	fi := mustParse(t, []byte(`f(a);`))
	call := firstCallExpr(t, fi)
	if got := findChildByType(call, "statement_block"); got != nil {
		t.Fatalf("expected nil, got %v", got.Type())
	}
}

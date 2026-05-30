//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestFindChildByType_Found 테스트
package express

import "testing"

func TestFindChildByType_Found(t *testing.T) {
	fi := mustParse(t, []byte(`f(a);`))
	call := firstCallExpr(t, fi)
	if got := findChildByType(call, "arguments"); got == nil {
		t.Fatal("expected arguments child")
	}
}

//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestIsResStatusCall_True 테스트
package express

import "testing"

func TestIsResStatusCall_True(t *testing.T) {
	fi := mustParse(t, []byte(`res.status(200);`))
	if !isResStatusCall(firstCallExpr(t, fi), fi.Src) {
		t.Fatal("expected true")
	}
}

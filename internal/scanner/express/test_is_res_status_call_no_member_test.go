//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestIsResStatusCall_NoMember 테스트
package express

import "testing"

func TestIsResStatusCall_NoMember(t *testing.T) {
	fi := mustParse(t, []byte(`status(200);`))
	if isResStatusCall(firstCallExpr(t, fi), fi.Src) {
		t.Fatal("expected false")
	}
}

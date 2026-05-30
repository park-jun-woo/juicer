//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestIsResStatusCall_NonResObject 테스트
package express

import "testing"

func TestIsResStatusCall_NonResObject(t *testing.T) {
	fi := mustParse(t, []byte(`other.status(200);`))
	if isResStatusCall(firstCallExpr(t, fi), fi.Src) {
		t.Fatal("expected false")
	}
}

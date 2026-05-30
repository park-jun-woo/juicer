//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestIsResStatusCall_PropMismatch 테스트
package express

import "testing"

func TestIsResStatusCall_PropMismatch(t *testing.T) {
	fi := mustParse(t, []byte(`res.json({});`))
	if isResStatusCall(firstCallExpr(t, fi), fi.Src) {
		t.Fatal("expected false")
	}
}

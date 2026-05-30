//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestIsResMethodCall_DirectRes 테스트
package express

import "testing"

func TestIsResMethodCall_DirectRes(t *testing.T) {
	fi := mustParse(t, []byte(`res.json({});`))
	m, ok := isResMethodCall(firstCallExpr(t, fi), fi.Src)
	if !ok || m != "json" {
		t.Fatalf("got %q,%v", m, ok)
	}
}

//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestIsResMethodCall_NonResObject 테스트
package express

import "testing"

func TestIsResMethodCall_NonResObject(t *testing.T) {
	fi := mustParse(t, []byte(`other.json({});`))
	if _, ok := isResMethodCall(firstCallExpr(t, fi), fi.Src); ok {
		t.Fatal("expected false for non-res object")
	}
}

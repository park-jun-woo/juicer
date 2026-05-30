//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestIsResMethodCall_ResNonMethod 테스트
package express

import "testing"

func TestIsResMethodCall_ResNonMethod(t *testing.T) {
	fi := mustParse(t, []byte(`res.render('v');`))
	if _, ok := isResMethodCall(firstCallExpr(t, fi), fi.Src); ok {
		t.Fatal("expected false for render")
	}
}

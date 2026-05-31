//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestIsResMethodCall_ResNonMethod 테스트
package express

import "testing"

// render/redirect는 Phase140부터 res 응답 메서드로 인식된다.
// res.cookie 등 비-응답 res 메서드는 여전히 false.
func TestIsResMethodCall_ResNonMethod(t *testing.T) {
	fi := mustParse(t, []byte(`res.cookie('k', 'v');`))
	if _, ok := isResMethodCall(firstCallExpr(t, fi), fi.Src); ok {
		t.Fatal("expected false for res.cookie")
	}
}

//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestIsExpressCall_OtherFn 테스트
package express

import "testing"

func TestIsExpressCall_OtherFn(t *testing.T) {
	fi := mustParse(t, []byte(`const a = foo();`))
	if isExpressCall(firstCallExpr(t, fi), fi.Src) {
		t.Fatal("expected false for foo()")
	}
}

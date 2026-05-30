//ff:func feature=scan type=test control=sequence topic=echo
//ff:what TestIsEchoContextType 테스트
package echo

import "testing"

func TestIsEchoContextType(t *testing.T) {
	if !isEchoContextType(parseExpr(t, "echo.Context")) {
		t.Fatal("expected true for echo.Context")
	}
	if isEchoContextType(parseExpr(t, "gin.Context")) {
		t.Fatal("gin.Context should be false")
	}
	if isEchoContextType(parseExpr(t, "Foo")) {
		t.Fatal("plain ident should be false")
	}
}

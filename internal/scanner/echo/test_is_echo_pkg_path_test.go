//ff:func feature=scan type=test control=sequence
//ff:what TestIsEchoPkgPath 테스트
package echo

import "testing"

func TestIsEchoPkgPath(t *testing.T) {
	if !isEchoPkgPath("github.com/labstack/echo/v4") {
		t.Fatal("expected v4 to be recognized")
	}
	if !isEchoPkgPath("github.com/labstack/echo/v5") {
		t.Fatal("expected v5 to be recognized")
	}
	if isEchoPkgPath("github.com/gin-gonic/gin") {
		t.Fatal("expected gin to be rejected")
	}
}

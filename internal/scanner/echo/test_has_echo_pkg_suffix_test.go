//ff:func feature=scan type=test control=sequence
//ff:what TestHasEchoPkgSuffix 테스트
package echo

import "testing"

func TestHasEchoPkgSuffix(t *testing.T) {
	if !hasEchoPkgSuffix("vendor/github.com/labstack/echo/v4") {
		t.Fatal("expected v4 suffix to match")
	}
	if !hasEchoPkgSuffix("github.com/labstack/echo/v5") {
		t.Fatal("expected v5 to match")
	}
	if hasEchoPkgSuffix("github.com/labstack/gommon") {
		t.Fatal("expected non-echo to be rejected")
	}
}

//ff:func feature=scan type=test control=sequence
//ff:what TestIsFiberRouterTypeInfo 테스트
package fiber

import "testing"

func TestIsFiberRouterTypeInfo(t *testing.T) {
	src := `package m
type App struct{}
type Other struct{}
var PA = &App{}     // *m.App — name "App" in set but pkg "m" not fiber
var VA = App{}      // m.App (non-pointer named)
var O  = &Other{}   // *Other — name not in router set
var I  = 5          // not pointer/named
`

	if isFiberRouterTypeInfo(typeOfVar(t, src, "PA")) {
		t.Error("local *App should be false (non-fiber pkg)")
	}

	if isFiberRouterTypeInfo(typeOfVar(t, src, "VA")) {
		t.Error("local App should be false")
	}

	if isFiberRouterTypeInfo(typeOfVar(t, src, "O")) {
		t.Error("*Other should be false")
	}

	if isFiberRouterTypeInfo(typeOfVar(t, src, "I")) {
		t.Error("int should be false")
	}
}

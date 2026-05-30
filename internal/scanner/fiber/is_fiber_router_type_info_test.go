//ff:func feature=scan type=test control=sequence
//ff:what isFiberRouterTypeInfo — types.Type 기반 라우터 타입 판정 테스트
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
	// *m.App: name in set but wrong package -> false
	if isFiberRouterTypeInfo(typeOfVar(t, src, "PA")) {
		t.Error("local *App should be false (non-fiber pkg)")
	}
	// non-pointer m.App -> named path, name in set, wrong pkg -> false
	if isFiberRouterTypeInfo(typeOfVar(t, src, "VA")) {
		t.Error("local App should be false")
	}
	// *Other: name not in set -> false
	if isFiberRouterTypeInfo(typeOfVar(t, src, "O")) {
		t.Error("*Other should be false")
	}
	// int -> not pointer/named -> false
	if isFiberRouterTypeInfo(typeOfVar(t, src, "I")) {
		t.Error("int should be false")
	}
}

func TestIsFiberRouterTypeInfo_ImportedNamed(t *testing.T) {
	// a named type from a real non-fiber package -> false
	src := `package m
import "bytes"
var B = bytes.Buffer{}
`
	if isFiberRouterTypeInfo(typeOfVar(t, src, "B")) {
		t.Error("bytes.Buffer should be false")
	}
}

//ff:func feature=scan type=test control=sequence
//ff:what TestIsFiberContextTypeInfo 테스트
package fiber

import "testing"

func TestIsFiberContextTypeInfo(t *testing.T) {
	src := `package m
type Ctx struct{}
type Other struct{}
var P = &Ctx{}      // pointer to named "Ctx" but pkg path is "m" (not fiber)
var O = &Other{}    // pointer to named "Other"
var I = 5           // not a pointer
var PB = new(int)   // pointer to non-named (basic)
`

	if isFiberContextTypeInfo(typeOfVar(t, src, "P")) {
		t.Error("local *Ctx (non-fiber pkg) should be false")
	}

	if isFiberContextTypeInfo(typeOfVar(t, src, "O")) {
		t.Error("*Other should be false")
	}

	if isFiberContextTypeInfo(typeOfVar(t, src, "I")) {
		t.Error("int should be false")
	}

	if isFiberContextTypeInfo(typeOfVar(t, src, "PB")) {
		t.Error("*int should be false")
	}
}

//ff:func feature=scan type=test control=sequence topic=echo
//ff:what TestIsEchoRouterTypeInfo_Round5 테스트
package echo

import (
	"go/types"
	"testing"
)

func TestIsEchoRouterTypeInfo_Round5(t *testing.T) {
	_, _, pkg := checkEchoPkg(t, `package echo
type Echo struct{}
type Group struct{}
type Other struct{}
`)
	echoT := pkg.Scope().Lookup("Echo").Type()
	if !isEchoRouterTypeInfo(echoT) {
		t.Fatal("Echo should be a router type")
	}

	grpT := pkg.Scope().Lookup("Group").Type()
	if !isEchoRouterTypeInfo(types.NewPointer(grpT)) {
		t.Fatal("*Group should be a router type")
	}
	otherT := pkg.Scope().Lookup("Other").Type()
	if isEchoRouterTypeInfo(otherT) {
		t.Fatal("Other should not be a router type")
	}
}

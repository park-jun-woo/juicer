//ff:func feature=scan type=test control=iteration dimension=1 topic=dotnet
//ff:what TestMethodAllowsBody_Round5 테스트
package dotnet

import "testing"

func TestMethodAllowsBody_Round5(t *testing.T) {
	for _, m := range []string{"POST", "PUT", "PATCH"} {
		if !methodAllowsBody(m) {
			t.Errorf("%s should allow body", m)
		}
	}
	for _, m := range []string{"GET", "DELETE", "HEAD", ""} {
		if methodAllowsBody(m) {
			t.Errorf("%s should not allow body", m)
		}
	}
}

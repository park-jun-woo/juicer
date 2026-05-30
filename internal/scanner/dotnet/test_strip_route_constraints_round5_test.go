//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what TestStripRouteConstraints_Round5 테스트
package dotnet

import "testing"

func TestStripRouteConstraints_Round5(t *testing.T) {
	if got := stripRouteConstraints("api/{id:int}"); got != "api/{id}" {
		t.Errorf("constraint: got %q", got)
	}
	if got := stripRouteConstraints("api/{*slug}"); got != "api/{slug}" {
		t.Errorf("catchall: got %q", got)
	}
	if got := stripRouteConstraints("api/plain"); got != "api/plain" {
		t.Errorf("plain: got %q", got)
	}
}

//ff:func feature=scan type=test topic=express control=sequence
//ff:what isRoutePathArg 경로 인자(* 또는 /시작) 판별 테스트
package express

import "testing"

func TestIsRoutePathArg(t *testing.T) {
	if !isRoutePathArg("*") {
		t.Error("* should be route path")
	}
	if !isRoutePathArg("/users") {
		t.Error("/users should be route path")
	}
	if isRoutePathArg("users") {
		t.Error("no leading slash should be false")
	}
	if isRoutePathArg("") {
		t.Error("empty should be false")
	}
}

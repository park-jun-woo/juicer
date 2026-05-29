//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what containsString: 존재/부재를 검증
package fastify

import "testing"

func TestContainsString(t *testing.T) {
	if !containsString([]string{"x", "y"}, "y") {
		t.Error("expected y present")
	}
	if containsString([]string{"x", "y"}, "z") {
		t.Error("expected z absent")
	}
}

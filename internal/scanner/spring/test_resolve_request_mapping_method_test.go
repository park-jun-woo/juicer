//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestResolveRequestMappingMethod 테스트
package spring

import "testing"

func TestResolveRequestMappingMethod(t *testing.T) {
	if got := resolveRequestMappingMethod("RequestMethod.POST"); got != "POST" {
		t.Fatalf("got %q", got)
	}
	if got := resolveRequestMappingMethod("{RequestMethod.GET}"); got != "GET" {
		t.Fatalf("braces: %q", got)
	}
	if got := resolveRequestMappingMethod("unknown"); got != "" {
		t.Fatalf("unknown: %q", got)
	}
}

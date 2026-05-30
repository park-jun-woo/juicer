//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestExtractGenericInner 테스트
package quarkus

import "testing"

func TestExtractGenericInner(t *testing.T) {
	if got := extractGenericInner("List<UserDto>"); got != "UserDto" {
		t.Fatalf("got %q", got)
	}
	if got := extractGenericInner("String"); got != "" {
		t.Fatalf("got %q", got)
	}
}

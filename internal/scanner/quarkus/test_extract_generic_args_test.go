//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestExtractGenericArgs 테스트
package quarkus

import "testing"

func TestExtractGenericArgs(t *testing.T) {
	if got := extractGenericArgs("Map<String, Long>"); got != "String, Long" {
		t.Fatalf("got %q", got)
	}
	if got := extractGenericArgs("String"); got != "" {
		t.Fatalf("got %q", got)
	}
}

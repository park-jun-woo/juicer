//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestExtractElementPairValue_Round5 테스트
package quarkus

import "testing"

func TestExtractElementPairValue_Round5(t *testing.T) {
	root, src := qParse(t, `@RolesAllowed(value = "admin") class C {}`)
	pair := qFirst(t, root, "element_value_pair")
	v, ok := extractElementPairValue(pair, src, "value")
	if !ok || v != "admin" {
		t.Fatalf("got %q %v", v, ok)
	}
	if _, ok := extractElementPairValue(pair, src, "other"); ok {
		t.Fatal("expected miss for wrong key")
	}
}

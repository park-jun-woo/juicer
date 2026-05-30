//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestExtractElementPairIntValue_Round5 테스트
package quarkus

import "testing"

func TestExtractElementPairIntValue_Round5(t *testing.T) {
	root, src := qParse(t, `@Max(status = 404) class C {}`)
	pair := qFirst(t, root, "element_value_pair")
	v, ok := extractElementPairIntValue(pair, src, "status")
	if !ok || v != 404 {
		t.Fatalf("got %d %v", v, ok)
	}
}

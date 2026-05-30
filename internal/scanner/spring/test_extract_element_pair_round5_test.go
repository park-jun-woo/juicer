//ff:func feature=scan type=test control=iteration dimension=1 topic=spring
//ff:what TestExtractElementPair_Round5 테스트
package spring

import "testing"

func TestExtractElementPair_Round5(t *testing.T) {
	root, src := sParse(t, `@Foo(name = "bar", code = 404) class C {}`)
	pairs := findAllByType(root, "element_value_pair")
	var sval string
	var ival int
	for _, p := range pairs {
		if v, ok := extractElementPairValue(p, src, "name"); ok {
			sval = v
		}
		if v, ok := extractElementPairIntValue(p, src, "code"); ok {
			ival = v
		}
	}
	if sval != "bar" {
		t.Errorf("name: %q", sval)
	}
	if ival != 404 {
		t.Errorf("code: %d", ival)
	}
}

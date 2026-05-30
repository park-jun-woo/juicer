//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestExtractRuleStrings_NoValue 테스트
package laravel

import "testing"

func TestExtractRuleStrings_NoValue(t *testing.T) {
	fi := mustParsePHP(t, `<?php $x = ['name'];`)
	elem := findAllByType(fi.root, "array_element_initializer")[0]
	if rules := extractRuleStrings(elem, fi.src); rules != nil {
		t.Fatalf("expected nil, got %v", rules)
	}
}

//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestExtractRuleStrings_Array 테스트
package laravel

import "testing"

func TestExtractRuleStrings_Array(t *testing.T) {
	fi := mustParsePHP(t, `<?php $x = ['role' => ['required', 'in:a,b']];`)
	elem := findAllByType(fi.root, "array_element_initializer")[0]
	rules := extractRuleStrings(elem, fi.src)
	if len(rules) != 2 {
		t.Fatalf("got %v", rules)
	}
}

//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestExtractRuleStrings_Pipe 테스트
package laravel

import "testing"

func TestExtractRuleStrings_Pipe(t *testing.T) {
	fi := mustParsePHP(t, `<?php $x = ['name' => 'required|string|max:255'];`)
	elem := findAllByType(fi.root, "array_element_initializer")[0]
	rules := extractRuleStrings(elem, fi.src)
	if len(rules) != 3 || rules[0] != "required" || rules[2] != "max:255" {
		t.Fatalf("got %v", rules)
	}
}

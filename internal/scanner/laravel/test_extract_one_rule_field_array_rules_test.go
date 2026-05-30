//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestExtractOneRuleField_ArrayRules 테스트
package laravel

import "testing"

func TestExtractOneRuleField_ArrayRules(t *testing.T) {
	fi := mustParsePHP(t, `<?php $x = ['role' => ['required', 'in:a,b']];`)
	elems := findAllByType(fi.root, "array_element_initializer")
	field := extractOneRuleField(elems[0], fi.src)
	if field == nil || field.Name != "role" {
		t.Fatalf("got %+v", field)
	}
	if len(field.Enum) != 2 {
		t.Fatalf("enum: %+v", field.Enum)
	}
}

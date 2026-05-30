//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestExtractFieldsFromRulesMethod_NoArray 테스트
package laravel

import "testing"

func TestExtractFieldsFromRulesMethod_NoArray(t *testing.T) {
	fi := mustParsePHP(t, `<?php class R { public function rules() { return null; } }`)
	method := findAllByType(fi.root, "method_declaration")[0]
	if got := extractFieldsFromRulesMethod(method, fi.src); got != nil {
		t.Fatalf("expected nil, got %+v", got)
	}
}

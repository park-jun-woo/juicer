//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestExtractFieldsFromRulesMethod 테스트
package laravel

import "testing"

func TestExtractFieldsFromRulesMethod(t *testing.T) {
	fi := mustParsePHP(t, `<?php class R { public function rules(): array {
		return [ 'name' => 'required|string', 'age' => 'integer' ];
	} }`)
	method := findAllByType(fi.root, "method_declaration")[0]
	fields := extractFieldsFromRulesMethod(method, fi.src)
	if len(fields) != 2 || fields[0].Name != "name" {
		t.Fatalf("got %+v", fields)
	}
}

//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestExtractRulesFromFile 테스트
package laravel

import "testing"

func TestExtractRulesFromFile(t *testing.T) {
	fi := mustParsePHP(t, `<?php class StoreReq {
		public function rules(): array { return [ 'name' => 'required|string' ]; }
	}`)
	fields := extractRulesFromFile(&fi, "StoreReq")
	if len(fields) != 1 || fields[0].Name != "name" {
		t.Fatalf("got %+v", fields)
	}
}

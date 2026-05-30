//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestExtractArrayKeys_NoReturnArray 테스트
package laravel

import "testing"

func TestExtractArrayKeys_NoReturnArray(t *testing.T) {
	fi := mustParsePHP(t, `<?php class R { public function toArray($r) { return null; } }`)
	methods := findAllByType(fi.root, "method_declaration")
	if got := extractArrayKeys(methods[0], fi.src); got != nil {
		t.Fatalf("expected nil, got %+v", got)
	}
}

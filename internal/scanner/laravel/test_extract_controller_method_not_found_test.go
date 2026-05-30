//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestExtractControllerMethod_NotFound 테스트
package laravel

import "testing"

func TestExtractControllerMethod_NotFound(t *testing.T) {
	fi := mustParsePHP(t, `<?php class C { public function index() {} }`)
	if cm := extractControllerMethod(&fi, "missing"); cm != nil {
		t.Fatalf("expected nil, got %+v", cm)
	}
}

//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestExtractControllerMethod 테스트
package laravel

import "testing"

func TestExtractControllerMethod(t *testing.T) {
	fi := mustParsePHP(t, `<?php class C { public function show(int $id) { return $id; } }`)
	cm := extractControllerMethod(&fi, "show")
	if cm == nil {
		t.Fatal("nil cm")
	}
	if cm.name != "show" || len(cm.params) != 1 || cm.params[0].name != "id" {
		t.Fatalf("got %+v", cm)
	}
	if len(cm.returnNodes) != 1 {
		t.Fatalf("returns: %+v", cm.returnNodes)
	}
}

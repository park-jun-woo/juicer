//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestFindAnyClassMethod 테스트
package laravel

import "testing"

func TestFindAnyClassMethod(t *testing.T) {
	fi := mustParsePHP(t, `<?php class A {} class B { public function bar() {} }`)
	if findAnyClassMethod(&fi, "bar") == nil {
		t.Fatal("expected to find bar in any class")
	}
	if findAnyClassMethod(&fi, "nope") != nil {
		t.Fatal("expected nil")
	}
}

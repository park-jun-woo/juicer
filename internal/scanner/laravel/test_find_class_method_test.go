//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestFindClassMethod 테스트
package laravel

import "testing"

func TestFindClassMethod(t *testing.T) {
	fi := mustParsePHP(t, `<?php class A { public function foo() {} } class B { public function bar() {} }`)
	if findClassMethod(&fi, "B", "bar") == nil {
		t.Fatal("expected to find B::bar")
	}
	if findClassMethod(&fi, "A", "bar") != nil {
		t.Fatal("A has no bar")
	}
	if findClassMethod(&fi, "Missing", "foo") != nil {
		t.Fatal("missing class")
	}
}

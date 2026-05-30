//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestFindMethodInClass 테스트
package laravel

import "testing"

func TestFindMethodInClass(t *testing.T) {
	fi := mustParsePHP(t, `<?php class C { public function foo() {} public function bar() {} }`)
	cls := findAllByType(fi.root, "class_declaration")[0]
	if findMethodInClass(cls, fi.src, "bar") == nil {
		t.Fatal("expected bar")
	}
	if findMethodInClass(cls, fi.src, "nope") != nil {
		t.Fatal("expected nil")
	}
}

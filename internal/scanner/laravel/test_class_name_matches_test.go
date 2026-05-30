//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestClassNameMatches 테스트
package laravel

import "testing"

func TestClassNameMatches(t *testing.T) {
	fi := mustParsePHP(t, `<?php class Foo {}`)
	cls := findAllByType(fi.root, "class_declaration")[0]
	if !classNameMatches(cls, fi.src, "Foo") {
		t.Fatal("expected match Foo")
	}
	if classNameMatches(cls, fi.src, "Bar") {
		t.Fatal("unexpected match Bar")
	}
}

//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestFindClassConstantAccess_Direct 테스트
package laravel

import "testing"

func TestFindClassConstantAccess_Direct(t *testing.T) {
	fi := mustParsePHP(t, `<?php $x = Foo::class;`)

	exprs := findAllByType(fi.root, "class_constant_access_expression")
	if len(exprs) == 0 {
		t.Fatal("no expr")
	}
	parent := exprs[0]
	if got := findClassConstantAccess(parent); got == nil {
		t.Fatal("expected to find self/child access")
	}
}

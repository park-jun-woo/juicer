//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestResourceTypeName_NotResource 테스트
package laravel

import "testing"

func TestResourceTypeName_NotResource(t *testing.T) {
	fi := mustParsePHP(t, `<?php $x = new Foo($u);`)
	ocs := findAllByType(fi.root, "object_creation_expression")
	if len(ocs) == 0 {
		t.Skip("no object creation")
	}
	if got := resourceTypeName(ocs[0], fi.src); got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}

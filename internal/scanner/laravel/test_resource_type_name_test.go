//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestResourceTypeName 테스트
package laravel

import "testing"

func TestResourceTypeName(t *testing.T) {
	fi := mustParsePHP(t, `<?php $x = new UserResource($u);`)
	ocs := findAllByType(fi.root, "object_creation_expression")
	if len(ocs) == 0 {
		t.Skip("no object creation")
	}
	if got := resourceTypeName(ocs[0], fi.src); got != "UserResource" {
		t.Fatalf("got %q", got)
	}
}

//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestCollectionResourceName 테스트
package laravel

import "testing"

func TestCollectionResourceName(t *testing.T) {
	fi := mustParsePHP(t, `<?php $x = UserResource::collection($users);`)
	scoped := findAllByType(fi.root, "scoped_call_expression")
	if len(scoped) == 0 {
		t.Skip("no scoped call")
	}
	if got := collectionResourceName(scoped[0], fi.src); got != "UserResource" {
		t.Fatalf("got %q", got)
	}
}

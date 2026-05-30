//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestMemberCallName 테스트
package laravel

import "testing"

func TestMemberCallName(t *testing.T) {
	fi := mustParsePHP(t, `<?php $obj->doThing();`)
	mcs := findAllByType(fi.root, "member_call_expression")
	if len(mcs) == 0 {
		t.Skip("no member call")
	}
	if got := memberCallName(mcs[0], fi.src); got != "doThing" {
		t.Fatalf("got %q", got)
	}
}

//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestSameNode 테스트
package laravel

import "testing"

func TestSameNode(t *testing.T) {
	fi := mustParsePHP(t, `<?php foo();`)
	calls := findAllByType(fi.root, "function_call_expression")
	if len(calls) == 0 {
		t.Skip("no call")
	}
	if !sameNode(calls[0], calls[0]) {
		t.Fatal("node should equal itself")
	}
	names := findAllByType(fi.root, "name")
	if len(names) > 0 && sameNode(calls[0], names[0]) {
		t.Fatal("different nodes should not be equal")
	}
}

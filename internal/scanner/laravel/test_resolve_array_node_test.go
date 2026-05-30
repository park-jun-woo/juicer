//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestResolveArrayNode 테스트
package laravel

import "testing"

func TestResolveArrayNode(t *testing.T) {

	fi := mustParsePHP(t, `<?php $x = [1, 2];`)
	arrs := findAllByType(fi.root, "array_creation_expression")
	if resolveArrayNode(arrs[0]) == nil {
		t.Fatal("self array node")
	}

	str := findAllByType(fi.root, "integer")
	if len(str) > 0 && resolveArrayNode(str[0]) != nil {
		t.Fatal("expected nil for non-array")
	}
}

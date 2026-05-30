//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestWalkChain_Other 테스트
package laravel

import "testing"

func TestWalkChain_Other(t *testing.T) {
	fi := mustParsePHP(t, `<?php $x = 1;`)
	ints := findAllByType(fi.root, "integer")
	if len(ints) == 0 {
		t.Skip("no int")
	}
	prefix := ""
	var mw []string
	walkChain(ints[0], fi, &prefix, &mw)
	if prefix != "" || mw != nil {
		t.Fatalf("unexpected: %q %v", prefix, mw)
	}
}

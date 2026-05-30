//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestExtractControllerAction_NotArray 테스트
package laravel

import "testing"

func TestExtractControllerAction_NotArray(t *testing.T) {
	fi := mustParsePHP(t, `<?php $x = 'closure';`)
	str := findAllByType(fi.root, "string")[0]
	ctrl, action := extractControllerAction(str, fi.src)
	if ctrl != "" || action != "" {
		t.Fatalf("got %q %q", ctrl, action)
	}
}

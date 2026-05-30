//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestExtractClassReference_NoAccess 테스트
package laravel

import "testing"

func TestExtractClassReference_NoAccess(t *testing.T) {
	fi := mustParsePHP(t, `<?php $x = 'plain';`)
	str := findAllByType(fi.root, "string")[0]
	if got := extractClassReference(str, fi.src); got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}

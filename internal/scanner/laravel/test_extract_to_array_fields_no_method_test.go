//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestExtractToArrayFields_NoMethod 테스트
package laravel

import "testing"

func TestExtractToArrayFields_NoMethod(t *testing.T) {
	fi := mustParsePHP(t, `<?php class R {}`)
	if got := extractToArrayFields(&fi, "R"); got != nil {
		t.Fatalf("expected nil, got %+v", got)
	}
}

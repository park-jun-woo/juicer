//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestFindChildByType 테스트
package laravel

import "testing"

func TestFindChildByType(t *testing.T) {
	fi := mustParsePHP(t, `<?php foo();`)
	calls := findAllByType(fi.root, "function_call_expression")
	if len(calls) == 0 {
		t.Skip("no call")
	}
	if findChildByType(calls[0], "arguments") == nil {
		t.Fatal("expected arguments")
	}
	if findChildByType(calls[0], "object") != nil {
		t.Fatal("expected nil for missing type")
	}
}

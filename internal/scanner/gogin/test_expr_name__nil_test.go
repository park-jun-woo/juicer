//ff:func feature=scan type=extract control=sequence
//ff:what TestExprName_Nil 테스트
package gogin

import "testing"

func TestExprName_Nil(t *testing.T) {
	got := exprName(nil)
	if got != "" {
		t.Fatalf("expected empty, got %s", got)
	}
}

//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestResolveViewName_Round5 테스트
package django

import "testing"

func TestResolveViewName_Round5(t *testing.T) {

	got := resolveViewName("MyView")
	if got == "" {
		t.Fatal("expected non-empty resolution")
	}
}

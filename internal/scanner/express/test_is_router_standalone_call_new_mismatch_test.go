//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestIsRouterStandaloneCall_NewMismatch 테스트
package express

import "testing"

func TestIsRouterStandaloneCall_NewMismatch(t *testing.T) {

	fi := mustParse(t, []byte(`const r = new Other();`))
	if isRouterStandaloneCall(findInitValue(firstDeclarator(t, fi)), fi.Src, map[string]bool{"Router": true}) {
		t.Fatal("expected false for new Other()")
	}
}

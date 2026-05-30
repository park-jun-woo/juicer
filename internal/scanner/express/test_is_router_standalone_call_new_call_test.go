//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestIsRouterStandaloneCall_NewCall 테스트
package express

import "testing"

func TestIsRouterStandaloneCall_NewCall(t *testing.T) {

	fi := mustParse(t, []byte(`const r = new Router();`))
	if !isRouterStandaloneCall(findInitValue(firstDeclarator(t, fi)), fi.Src, map[string]bool{"Router": true}) {
		t.Fatal("expected true for new Router()")
	}
}

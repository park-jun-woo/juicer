//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestIsRouterStandaloneCall_PlainCall 테스트
package express

import "testing"

func TestIsRouterStandaloneCall_PlainCall(t *testing.T) {
	fi := mustParse(t, []byte(`const r = Router();`))
	if !isRouterStandaloneCall(findInitValue(firstDeclarator(t, fi)), fi.Src, map[string]bool{"Router": true}) {
		t.Fatal("expected true for Router()")
	}
}

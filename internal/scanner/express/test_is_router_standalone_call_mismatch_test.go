//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestIsRouterStandaloneCall_Mismatch 테스트
package express

import "testing"

func TestIsRouterStandaloneCall_Mismatch(t *testing.T) {
	fi := mustParse(t, []byte(`const r = Other();`))
	if isRouterStandaloneCall(findInitValue(firstDeclarator(t, fi)), fi.Src, map[string]bool{"Router": true}) {
		t.Fatal("expected false for Other()")
	}
}

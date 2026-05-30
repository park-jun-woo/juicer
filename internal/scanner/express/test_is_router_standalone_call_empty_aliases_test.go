//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestIsRouterStandaloneCall_EmptyAliases 테스트
package express

import "testing"

func TestIsRouterStandaloneCall_EmptyAliases(t *testing.T) {
	fi := mustParse(t, []byte(`const r = Router();`))
	if isRouterStandaloneCall(findInitValue(firstDeclarator(t, fi)), fi.Src, map[string]bool{}) {
		t.Fatal("expected false for empty aliases")
	}
}

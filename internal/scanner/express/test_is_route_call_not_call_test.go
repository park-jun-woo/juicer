//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestIsRouteCall_NotCall 테스트
package express

import "testing"

func TestIsRouteCall_NotCall(t *testing.T) {
	fi := mustParse(t, []byte(`router;`))
	ids := findAllByType(fi.Root, "identifier")
	if isRouteCall(ids[0], fi.Src, map[string]bool{"router": true}) {
		t.Fatal("expected false")
	}
}

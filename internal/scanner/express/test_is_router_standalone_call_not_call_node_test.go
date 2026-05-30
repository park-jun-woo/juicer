//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestIsRouterStandaloneCall_NotCallNode 테스트
package express

import "testing"

func TestIsRouterStandaloneCall_NotCallNode(t *testing.T) {

	fi := mustParse(t, []byte(`Router;`))
	ids := findAllByType(fi.Root, "identifier")
	if isRouterStandaloneCall(ids[0], fi.Src, map[string]bool{"Router": true}) {
		t.Fatal("expected false for identifier node")
	}
}

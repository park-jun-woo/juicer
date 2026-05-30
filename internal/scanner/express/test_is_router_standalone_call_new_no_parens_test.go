//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestIsRouterStandaloneCall_NewNoParens 테스트
package express

import "testing"

func TestIsRouterStandaloneCall_NewNoParens(t *testing.T) {

	fi := mustParse(t, []byte(`const r = new Router;`))
	v := findInitValue(firstDeclarator(t, fi))
	if v == nil {

		ne := findAllByType(fi.Root, "new_expression")
		if len(ne) == 0 {
			t.Skip("no new_expression produced")
		}
		v = ne[0]
	}
	if !isRouterStandaloneCall(v, fi.Src, map[string]bool{"Router": true}) {
		t.Fatal("expected true for new Router")
	}
}

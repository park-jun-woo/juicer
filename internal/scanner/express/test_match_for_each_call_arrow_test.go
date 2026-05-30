//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestMatchForEachCall_Arrow 테스트
package express

import "testing"

func TestMatchForEachCall_Arrow(t *testing.T) {
	fi := mustParse(t, []byte(`routes.forEach(r => { parent.use('/x', r); });`))
	if got := matchForEachCall(firstCallExpr(t, fi), fi.Src, map[string]bool{"parent": true}); got != "routes" {
		t.Fatalf("got %q", got)
	}
}

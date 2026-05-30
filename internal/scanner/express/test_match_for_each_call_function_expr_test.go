//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestMatchForEachCall_FunctionExpr 테스트
package express

import "testing"

func TestMatchForEachCall_FunctionExpr(t *testing.T) {
	fi := mustParse(t, []byte(`routes.forEach(function(r) { parent.use('/x', r); });`))
	if got := matchForEachCall(firstCallExpr(t, fi), fi.Src, map[string]bool{"parent": true}); got != "routes" {
		t.Fatalf("got %q", got)
	}
}

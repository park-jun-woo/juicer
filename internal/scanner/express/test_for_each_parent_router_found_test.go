//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestForEachParentRouter_Found 테스트
package express

import "testing"

func TestForEachParentRouter_Found(t *testing.T) {
	fi := mustParse(t, []byte(`arr.forEach(r => { parent.use('/x', r); });`))
	call := firstCallExpr(t, fi)
	got := forEachParentRouter(call, fi.Src, map[string]bool{"parent": true})
	if got != "parent" {
		t.Fatalf("got %q", got)
	}
}

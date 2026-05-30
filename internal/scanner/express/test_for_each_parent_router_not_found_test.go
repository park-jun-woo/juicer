//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestForEachParentRouter_NotFound 테스트
package express

import "testing"

func TestForEachParentRouter_NotFound(t *testing.T) {
	fi := mustParse(t, []byte(`arr.forEach(r => { doStuff(r); });`))
	call := firstCallExpr(t, fi)
	if got := forEachParentRouter(call, fi.Src, map[string]bool{"parent": true}); got != "" {
		t.Fatalf("got %q", got)
	}
}

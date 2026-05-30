//ff:func feature=scan type=test control=sequence topic=express
//ff:what forEachParentRouter: router.use 발견 / 미발견 분기
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

func TestForEachParentRouter_NotFound(t *testing.T) {
	fi := mustParse(t, []byte(`arr.forEach(r => { doStuff(r); });`))
	call := firstCallExpr(t, fi)
	if got := forEachParentRouter(call, fi.Src, map[string]bool{"parent": true}); got != "" {
		t.Fatalf("got %q", got)
	}
}

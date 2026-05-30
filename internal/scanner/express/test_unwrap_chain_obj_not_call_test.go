//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestUnwrapChain_ObjNotCall 테스트
package express

import "testing"

func TestUnwrapChain_ObjNotCall(t *testing.T) {

	fi := mustParse(t, []byte(`router.get('/x');`))
	if p, _, m := unwrapChain(firstCallExpr(t, fi), fi.Src, map[string]bool{"router": true}); p != "" || m != nil {
		t.Fatalf("expected empty, got p=%q m=%v", p, m)
	}
}

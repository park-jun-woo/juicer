//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestUnwrapChain_NoMember 테스트
package express

import "testing"

func TestUnwrapChain_NoMember(t *testing.T) {
	fi := mustParse(t, []byte(`foo();`))
	if p, _, m := unwrapChain(firstCallExpr(t, fi), fi.Src, nil); p != "" || m != nil {
		t.Fatalf("expected empty")
	}
}

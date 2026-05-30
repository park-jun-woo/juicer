//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestMatchForEachCall_NoMember 테스트
package express

import "testing"

func TestMatchForEachCall_NoMember(t *testing.T) {
	fi := mustParse(t, []byte(`forEach(cb);`))
	if got := matchForEachCall(firstCallExpr(t, fi), fi.Src, nil); got != "" {
		t.Fatalf("got %q", got)
	}
}

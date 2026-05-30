//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestMatchForEachCall_NotForEach 테스트
package express

import "testing"

func TestMatchForEachCall_NotForEach(t *testing.T) {
	fi := mustParse(t, []byte(`routes.map(r => r);`))
	if got := matchForEachCall(firstCallExpr(t, fi), fi.Src, nil); got != "" {
		t.Fatalf("got %q", got)
	}
}

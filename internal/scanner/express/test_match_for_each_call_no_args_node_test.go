//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestMatchForEachCall_NoArgsNode 테스트
package express

import "testing"

func TestMatchForEachCall_NoArgsNode(t *testing.T) {

	fi := mustParse(t, []byte("routes.forEach`x`;"))
	if got := matchForEachCall(firstCallExpr(t, fi), fi.Src, nil); got != "" {
		t.Fatalf("got %q", got)
	}
}

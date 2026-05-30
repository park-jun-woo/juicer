//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestMatchForEachCall_NoArrowOrFunc 테스트
package express

import "testing"

func TestMatchForEachCall_NoArrowOrFunc(t *testing.T) {
	fi := mustParse(t, []byte(`routes.forEach(cbRef);`))
	if got := matchForEachCall(firstCallExpr(t, fi), fi.Src, map[string]bool{"parent": true}); got != "" {
		t.Fatalf("got %q", got)
	}
}

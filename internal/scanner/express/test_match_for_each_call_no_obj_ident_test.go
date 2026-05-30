//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestMatchForEachCall_NoObjIdent 테스트
package express

import "testing"

func TestMatchForEachCall_NoObjIdent(t *testing.T) {

	fi := mustParse(t, []byte(`a.b.forEach(r => r);`))
	if got := matchForEachCall(firstCallExpr(t, fi), fi.Src, nil); got != "" {
		t.Fatalf("got %q", got)
	}
}

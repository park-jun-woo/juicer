//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestExtractValidateRequest_NoFnIdentifier 테스트
package express

import "testing"

func TestExtractValidateRequest_NoFnIdentifier(t *testing.T) {

	fi := mustParse(t, []byte(`mw.validate({ body: s });`))
	if got := extractValidateRequest(firstCallExpr(t, fi), fi.Src); got != nil {
		t.Fatalf("got %+v", got)
	}
}

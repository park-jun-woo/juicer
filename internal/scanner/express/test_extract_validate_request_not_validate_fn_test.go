//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestExtractValidateRequest_NotValidateFn 테스트
package express

import "testing"

func TestExtractValidateRequest_NotValidateFn(t *testing.T) {
	fi := mustParse(t, []byte(`foo({ body: s });`))
	if got := extractValidateRequest(firstCallExpr(t, fi), fi.Src); got != nil {
		t.Fatalf("got %+v", got)
	}
}

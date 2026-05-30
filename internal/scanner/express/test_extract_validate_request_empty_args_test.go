//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestExtractValidateRequest_EmptyArgs 테스트
package express

import "testing"

func TestExtractValidateRequest_EmptyArgs(t *testing.T) {
	fi := mustParse(t, []byte(`validate();`))
	if got := extractValidateRequest(firstCallExpr(t, fi), fi.Src); got != nil {
		t.Fatalf("got %+v", got)
	}
}

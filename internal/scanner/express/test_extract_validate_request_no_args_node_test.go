//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestExtractValidateRequest_NoArgsNode 테스트
package express

import "testing"

func TestExtractValidateRequest_NoArgsNode(t *testing.T) {
	fi := mustParse(t, []byte("validate`x`;"))
	if got := extractValidateRequest(firstCallExpr(t, fi), fi.Src); got != nil {
		t.Fatalf("got %+v", got)
	}
}

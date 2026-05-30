//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestExtractValidateRequest_NotCall 테스트
package express

import "testing"

func TestExtractValidateRequest_NotCall(t *testing.T) {
	fi := mustParse(t, []byte(`const x = 1;`))
	if got := extractValidateRequest(fi.Root, fi.Src); got != nil {
		t.Fatalf("got %+v", got)
	}
}

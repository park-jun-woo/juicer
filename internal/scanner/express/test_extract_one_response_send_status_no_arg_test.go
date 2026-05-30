//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestExtractOneResponse_SendStatusNoArg 테스트
package express

import "testing"

func TestExtractOneResponse_SendStatusNoArg(t *testing.T) {

	fi := mustParse(t, []byte(`res.sendStatus();`))
	r := extractOneResponse(firstCallExpr(t, fi), fi.Src)
	if r == nil || r.Status != "200" || r.Kind != "empty" {
		t.Fatalf("got %+v", r)
	}
}

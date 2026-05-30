//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestExtractOneResponse_SendDefault 테스트
package express

import "testing"

func TestExtractOneResponse_SendDefault(t *testing.T) {
	fi := mustParse(t, []byte(`res.send('ok');`))
	r := extractOneResponse(firstCallExpr(t, fi), fi.Src)
	if r == nil || r.Status != "200" || r.Kind != "text" {
		t.Fatalf("got %+v", r)
	}
}

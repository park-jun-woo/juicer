//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestExtractOneResponse_SendStatus 테스트
package express

import "testing"

func TestExtractOneResponse_SendStatus(t *testing.T) {
	fi := mustParse(t, []byte(`res.sendStatus(204);`))
	r := extractOneResponse(firstCallExpr(t, fi), fi.Src)
	if r == nil || r.Status != "204" || r.Kind != "empty" {
		t.Fatalf("got %+v", r)
	}
}

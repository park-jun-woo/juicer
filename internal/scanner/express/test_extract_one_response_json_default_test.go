//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestExtractOneResponse_JsonDefault 테스트
package express

import "testing"

func TestExtractOneResponse_JsonDefault(t *testing.T) {
	fi := mustParse(t, []byte(`res.json({a:1});`))
	r := extractOneResponse(firstCallExpr(t, fi), fi.Src)
	if r == nil || r.Status != "200" || r.Kind != "json" {
		t.Fatalf("got %+v", r)
	}
}

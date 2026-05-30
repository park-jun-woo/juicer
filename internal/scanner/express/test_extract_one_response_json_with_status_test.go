//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestExtractOneResponse_JsonWithStatus 테스트
package express

import "testing"

func TestExtractOneResponse_JsonWithStatus(t *testing.T) {
	fi := mustParse(t, []byte(`res.status(201).json({a:1});`))
	r := extractOneResponse(firstCallExpr(t, fi), fi.Src)
	if r == nil || r.Status != "201" || r.Kind != "json" {
		t.Fatalf("got %+v", r)
	}
}

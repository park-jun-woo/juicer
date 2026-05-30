//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestExtractOneResponse_OtherResMethod 테스트
package express

import "testing"

func TestExtractOneResponse_OtherResMethod(t *testing.T) {

	fi := mustParse(t, []byte(`res.render('view');`))
	r := extractOneResponse(firstCallExpr(t, fi), fi.Src)
	if r != nil {
		t.Fatalf("expected nil for render, got %+v", r)
	}
}

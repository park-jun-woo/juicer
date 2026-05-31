//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestExtractOneResponse_OtherResMethod 테스트
package express

import "testing"

// res.render는 Phase140부터 200/html로 인식된다 (이전: nil).
// res.cookie 등 비-응답 res 메서드는 여전히 nil.
func TestExtractOneResponse_OtherResMethod(t *testing.T) {
	fi := mustParse(t, []byte(`res.cookie('k', 'v');`))
	r := extractOneResponse(firstCallExpr(t, fi), fi.Src)
	if r != nil {
		t.Fatalf("expected nil for res.cookie, got %+v", r)
	}
}

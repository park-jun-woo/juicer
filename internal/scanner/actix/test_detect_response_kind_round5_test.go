//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestDetectResponseKind_Round5 테스트
package actix

import "testing"

func TestDetectResponseKind_Round5(t *testing.T) {
	fi := aFi(t, `async fn h() -> impl Responder { HttpResponse::Ok().json(user) }`)
	sid := aFirst(t, fi.root, "scoped_identifier")
	_ = detectResponseKind(sid, fi.src)
}

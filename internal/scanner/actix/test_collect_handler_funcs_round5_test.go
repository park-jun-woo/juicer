//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestCollectHandlerFuncs_Round5 테스트
package actix

import "testing"

func TestCollectHandlerFuncs_Round5(t *testing.T) {
	fi := aFi(t, `async fn list_users() -> impl Responder { "" }`)
	index := map[string]*handlerInfo{}
	collectHandlerFuncs(fi, index)
	if _, ok := index["list_users"]; !ok {
		t.Fatalf("handler not collected: %v", index)
	}
}

//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestFindHandlerBody_NamedFunction 테스트
package express

import "testing"

func TestFindHandlerBody_NamedFunction(t *testing.T) {
	fi := mustParse(t, []byte(`function getUsers(req, res) { res.json({}); }`))
	if body := findHandlerBody(fi, routeInfo{Handler: "getUsers"}); body == nil {
		t.Fatal("expected named function body")
	}
}

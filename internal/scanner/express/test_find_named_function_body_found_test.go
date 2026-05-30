//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestFindNamedFunctionBody_Found 테스트
package express

import "testing"

func TestFindNamedFunctionBody_Found(t *testing.T) {
	fi := mustParse(t, []byte(`function getUsers(req, res) { res.json({}); }`))
	if body := findNamedFunctionBody(fi, "getUsers"); body == nil {
		t.Fatal("expected body")
	}
}

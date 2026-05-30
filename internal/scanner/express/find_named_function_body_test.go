//ff:func feature=scan type=test control=sequence topic=express
//ff:what findNamedFunctionBody: 매칭 본문 반환 / 미발견 nil
package express

import "testing"

func TestFindNamedFunctionBody_Found(t *testing.T) {
	fi := mustParse(t, []byte(`function getUsers(req, res) { res.json({}); }`))
	if body := findNamedFunctionBody(fi, "getUsers"); body == nil {
		t.Fatal("expected body")
	}
}

func TestFindNamedFunctionBody_NotFound(t *testing.T) {
	fi := mustParse(t, []byte(`function getUsers() {}`))
	if body := findNamedFunctionBody(fi, "missing"); body != nil {
		t.Fatalf("expected nil, got %v", body.Type())
	}
}

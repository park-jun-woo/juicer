//ff:func feature=scan type=test control=sequence topic=express
//ff:what findHandlerBody: HandlerNode본문 / anonymous nil / 명명함수 조회 분기
package express

import "testing"

func TestFindHandlerBody_FromHandlerNode(t *testing.T) {
	fi := mustParse(t, []byte(`const h = (req, res) => { res.json({}); };`))
	ri := routeInfo{HandlerNode: firstArrow(t, fi)}
	if body := findHandlerBody(fi, ri); body == nil {
		t.Fatal("expected body from handler node")
	}
}

func TestFindHandlerBody_Anonymous(t *testing.T) {
	fi := mustParse(t, []byte(`const x = 1;`))
	if body := findHandlerBody(fi, routeInfo{Handler: "(anonymous)"}); body != nil {
		t.Fatalf("expected nil, got %v", body.Type())
	}
}

func TestFindHandlerBody_NamedFunction(t *testing.T) {
	fi := mustParse(t, []byte(`function getUsers(req, res) { res.json({}); }`))
	if body := findHandlerBody(fi, routeInfo{Handler: "getUsers"}); body == nil {
		t.Fatal("expected named function body")
	}
}

func TestFindHandlerBody_NamedNotFound(t *testing.T) {
	fi := mustParse(t, []byte(`const x = 1;`))
	if body := findHandlerBody(fi, routeInfo{Handler: "missing"}); body != nil {
		t.Fatalf("expected nil, got %v", body.Type())
	}
}

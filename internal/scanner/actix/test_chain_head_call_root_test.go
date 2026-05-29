//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestChainHeadCallRoot — 체인 헤드 식별(cfg vs web::scope, 인자 내부 무시)을 검증
package actix

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
)

func TestChainHeadCallRoot(t *testing.T) {
	src := []byte(`
fn config(cfg: &mut web::ServiceConfig) {
    cfg.service(web::resource("/a").route(web::get().to(a)));
    web::scope("/p").service(web::resource("/b"));
}
`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	// Collect the receiver of each outermost .service() call and check its head.
	heads := map[string]bool{}
	var walk func(n *sitter.Node)
	walk = func(n *sitter.Node) {
		if isTopLevelServiceCall(n, src) {
			fe := findChildByType(n, "field_expression")
			recv := findFieldReceiver(fe)
			heads[chainHeadCallRoot(recv, src)] = true
		}
		for i := 0; i < int(n.ChildCount()); i++ {
			walk(n.Child(i))
		}
	}
	walk(root)
	// cfg.service(...) is top-level: its receiver chain head is "cfg", even
	// though web::resource appears inside the argument.
	if !heads["cfg"] {
		t.Errorf("expected top-level cfg.service head 'cfg', got heads=%v", heads)
	}
	// web::scope("/p").service(...) is NOT top-level (skipped), so its receiver
	// head must never be recorded here.
	if heads["web::scope"] {
		t.Error("web::scope-rooted .service() must not be treated as top-level")
	}
}

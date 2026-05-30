//ff:func feature=scan type=test control=sequence topic=actix
//ff:what findRouteCalls — .route() 호출 인자 콜백 전달을 검증
package actix

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
)

func TestFindRouteCalls(t *testing.T) {
	src := []byte(`
fn f() {
    web::resource("/x")
        .route(web::get().to(a))
        .route(web::post().to(b));
}
`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	// findRouteCalls walks down a method chain from its outermost call, so we
	// pass the outermost call_expression of the chain.
	chain := findCallByFuncSuffix(root, src, ".route")
	if chain == nil {
		t.Fatal("no .route call found")
	}
	count := 0
	findRouteCalls(chain, src, func(arg *sitter.Node) {
		count++
	})
	if count != 2 {
		t.Fatalf("expected 2 route calls, got %d", count)
	}
}

func TestFindRouteCalls_None(t *testing.T) {
	// A chain with .service() but no .route() yields no callbacks.
	src := []byte(`fn f() { web::scope("/x").service(a); }`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	chain := findCallByFuncSuffix(root, src, ".service")
	if chain == nil {
		t.Fatal("no .service call found")
	}
	count := 0
	findRouteCalls(chain, src, func(arg *sitter.Node) { count++ })
	if count != 0 {
		t.Fatalf("expected 0 route calls, got %d", count)
	}
}

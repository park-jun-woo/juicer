//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestFindRouteCalls 테스트
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
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

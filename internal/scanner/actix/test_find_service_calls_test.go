//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestFindServiceCalls 테스트
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func TestFindServiceCalls(t *testing.T) {
	src := []byte(`
fn f() {
    web::scope("/x")
        .service(a)
        .service(b);
}
`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	chain := findCallByFuncSuffix(root, src, ".service")
	if chain == nil {
		t.Fatal("no .service call found")
	}
	count := 0
	findServiceCalls(chain, src, func(arg *sitter.Node) { count++ })
	if count != 2 {
		t.Fatalf("expected 2 service calls, got %d", count)
	}
}

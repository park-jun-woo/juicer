//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestFindRouteCalls_None 테스트
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func TestFindRouteCalls_None(t *testing.T) {

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

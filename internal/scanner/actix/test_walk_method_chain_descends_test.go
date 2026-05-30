//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestWalkMethodChain_Descends 테스트
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func TestWalkMethodChain_Descends(t *testing.T) {
	src := []byte(`fn f() { web::scope("/p").service(a).service(b); }`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	chain := findCallByFuncSuffix(root, src, ".service")
	if chain == nil {
		t.Fatal("no .service chain")
	}
	count := 0
	walkMethodChain(chain, src, "service", func(args *sitter.Node) { count++ })
	if count != 2 {
		t.Fatalf("expected 2 service callbacks, got %d", count)
	}
}

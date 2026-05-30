//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestWalkMethodChain_NoFieldExpr 테스트
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func TestWalkMethodChain_NoFieldExpr(t *testing.T) {

	src := []byte(`fn f() { web::get(); }`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	call := findCallByFuncSuffix(root, src, "::get")
	if call == nil {
		t.Fatal("no get call")
	}
	count := 0
	walkMethodChain(call, src, "service", func(args *sitter.Node) { count++ })
	if count != 0 {
		t.Fatalf("expected 0 callbacks, got %d", count)
	}
}

//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestExtractResponses_NoBlock 테스트
package actix

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestExtractResponses_NoBlock(t *testing.T) {

	src := []byte(`
trait T {
    async fn handler() -> HttpResponse;
}
`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	fn := findFuncByName(root, src, "handler")
	if fn == nil {
		t.Fatal("function signature not found")
	}
	idx := buildStructIndex([]*fileInfo{{src: src, root: root}})
	if resps := extractResponses(fn, src, idx, map[string][]scanner.Field{}); resps != nil {
		t.Fatalf("expected nil responses for block-less function, got %+v", resps)
	}
}

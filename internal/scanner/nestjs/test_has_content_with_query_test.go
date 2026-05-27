//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestHasContent_WithQuery 테스트
package nestjs

import (
	"testing"
	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestHasContent_WithQuery(t *testing.T) {
	r := &scanner.Request{Query: []scanner.Param{{Name: "page"}}}
	if !hasContent(r) {
		t.Fatal("expected true with query")
	}
}

//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestHasContent_WithPathParam 테스트
package nestjs

import (
	"testing"
	"github.com/park-jun-woo/juicer/internal/scanner"
)

func TestHasContent_WithPathParam(t *testing.T) {
	r := &scanner.Request{PathParams: []scanner.Param{{Name: "id"}}}
	if !hasContent(r) {
		t.Fatal("expected true with path param")
	}
}

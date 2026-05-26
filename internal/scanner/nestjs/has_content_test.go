//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestHasContent_Empty 테스트
package nestjs

import (
	"testing"
	"github.com/park-jun-woo/juicer/internal/scanner"
)

func TestHasContent_Empty(t *testing.T) {
	r := &scanner.Request{}
	if hasContent(r) {
		t.Fatal("expected false for empty request")
	}
}

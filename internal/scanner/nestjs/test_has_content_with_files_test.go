//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestHasContent_WithFiles 테스트
package nestjs

import (
	"testing"
	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestHasContent_WithFiles(t *testing.T) {
	r := &scanner.Request{Files: []scanner.Param{{Name: "f"}}}
	if !hasContent(r) {
		t.Fatal("expected true with files")
	}
}

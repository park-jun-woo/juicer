//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestHasContent_WithBody 테스트
package nestjs

import (
	"testing"
	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestHasContent_WithBody(t *testing.T) {
	r := &scanner.Request{Body: &scanner.Body{TypeName: "Dto"}}
	if !hasContent(r) {
		t.Fatal("expected true with body")
	}
}

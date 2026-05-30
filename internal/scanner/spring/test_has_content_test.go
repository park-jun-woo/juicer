//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestHasContent 테스트
package spring

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestHasContent(t *testing.T) {
	if hasContent(&scanner.Request{}) {
		t.Fatal("empty")
	}
	if !hasContent(&scanner.Request{Body: &scanner.Body{}}) {
		t.Fatal("body")
	}
}

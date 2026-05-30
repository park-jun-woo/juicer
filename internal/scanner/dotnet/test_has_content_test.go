//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what TestHasContent 테스트
package dotnet

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

//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what hasContent 테스트
package fastapi

import (
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestHasContent(t *testing.T) {
	if hasContent(&scanner.Request{}) {
		t.Fatal("empty request should not have content")
	}
	if !hasContent(&scanner.Request{PathParams: []scanner.Param{{Name: "id"}}}) {
		t.Fatal("request with path params should have content")
	}
	if !hasContent(&scanner.Request{Query: []scanner.Param{{Name: "q"}}}) {
		t.Fatal("request with query should have content")
	}
	if !hasContent(&scanner.Request{Body: &scanner.Body{}}) {
		t.Fatal("request with body should have content")
	}
	if !hasContent(&scanner.Request{Files: []scanner.Param{{Name: "f"}}}) {
		t.Fatal("request with files should have content")
	}
}

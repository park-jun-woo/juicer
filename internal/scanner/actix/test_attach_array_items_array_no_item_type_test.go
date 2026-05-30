//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestAttachArrayItems_ArrayNoItemType 테스트
package actix

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestAttachArrayItems_ArrayNoItemType(t *testing.T) {
	f := &scanner.Field{Name: "x", Type: "array"}
	attachArrayItems(f, openAPIType{Type: "array", Items: ""})
	if f.Fields != nil {
		t.Fatalf("expected no items when Items empty, got %+v", f.Fields)
	}
}

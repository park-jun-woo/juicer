//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestAttachArrayItems_NotArray 테스트
package actix

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestAttachArrayItems_NotArray(t *testing.T) {
	f := &scanner.Field{Name: "x", Type: "string"}
	attachArrayItems(f, openAPIType{Type: "string"})
	if f.Fields != nil {
		t.Fatalf("expected no items for non-array, got %+v", f.Fields)
	}
}

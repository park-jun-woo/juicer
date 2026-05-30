//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestAttachArrayItems_Array 테스트
package actix

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestAttachArrayItems_Array(t *testing.T) {
	f := &scanner.Field{Name: "tags", Type: "array"}
	attachArrayItems(f, openAPIType{Type: "array", Items: "String"})
	if len(f.Fields) != 1 {
		t.Fatalf("expected 1 items field, got %d", len(f.Fields))
	}
	if f.Fields[0].Name != "items" || f.Fields[0].Type != "string" {
		t.Errorf("unexpected items field: %+v", f.Fields[0])
	}
}

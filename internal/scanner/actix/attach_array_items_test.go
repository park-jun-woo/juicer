//ff:func feature=scan type=test control=sequence topic=actix
//ff:what attachArrayItems — 배열 타입 필드에 items 부착 분기를 검증
package actix

import (
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
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

func TestAttachArrayItems_NotArray(t *testing.T) {
	f := &scanner.Field{Name: "x", Type: "string"}
	attachArrayItems(f, openAPIType{Type: "string"})
	if f.Fields != nil {
		t.Fatalf("expected no items for non-array, got %+v", f.Fields)
	}
}

func TestAttachArrayItems_ArrayNoItemType(t *testing.T) {
	f := &scanner.Field{Name: "x", Type: "array"}
	attachArrayItems(f, openAPIType{Type: "array", Items: ""})
	if f.Fields != nil {
		t.Fatalf("expected no items when Items empty, got %+v", f.Fields)
	}
}

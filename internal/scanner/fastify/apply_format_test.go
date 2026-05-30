//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what applyFormat 테스트
package fastify

import (
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestApplyFormat_StringWithFormat(t *testing.T) {
	obj, src := firstObject(t, `{ type: "string", format: "email" }`)
	f := &scanner.Field{Type: "string"}
	applyFormat(f, obj, src)
	if f.Type != "email" {
		t.Fatalf("expected email, got %q", f.Type)
	}
}

func TestApplyFormat_NoFormat(t *testing.T) {
	obj, src := firstObject(t, `{ type: "string" }`)
	f := &scanner.Field{Type: "string"}
	applyFormat(f, obj, src)
	if f.Type != "string" {
		t.Fatalf("expected unchanged string, got %q", f.Type)
	}
}

func TestApplyFormat_NonStringType(t *testing.T) {
	// format present but type is not "string" -> unchanged
	obj, src := firstObject(t, `{ type: "integer", format: "int64" }`)
	f := &scanner.Field{Type: "integer"}
	applyFormat(f, obj, src)
	if f.Type != "integer" {
		t.Fatalf("expected unchanged integer, got %q", f.Type)
	}
}

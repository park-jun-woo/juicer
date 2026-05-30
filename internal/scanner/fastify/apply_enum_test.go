//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what applyEnum 테스트
package fastify

import (
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestApplyEnum_Array(t *testing.T) {
	obj, src := firstObject(t, `{ type: "string", enum: ["a", "b", "c"] }`)
	f := &scanner.Field{}
	applyEnum(f, obj, src)
	if len(f.Enum) != 3 || f.Enum[0] != "a" || f.Enum[2] != "c" {
		t.Fatalf("expected [a b c], got %v", f.Enum)
	}
}

func TestApplyEnum_NoEnum(t *testing.T) {
	obj, src := firstObject(t, `{ type: "string" }`)
	f := &scanner.Field{}
	applyEnum(f, obj, src)
	if len(f.Enum) != 0 {
		t.Fatalf("expected no enum, got %v", f.Enum)
	}
}

func TestApplyEnum_NotArray(t *testing.T) {
	// enum value is not an array literal -> early return
	obj, src := firstObject(t, `{ type: "string", enum: "x" }`)
	f := &scanner.Field{}
	applyEnum(f, obj, src)
	if len(f.Enum) != 0 {
		t.Fatalf("expected no enum when not array, got %v", f.Enum)
	}
}

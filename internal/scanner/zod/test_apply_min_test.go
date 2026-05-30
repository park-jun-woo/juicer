//ff:func feature=scan type=test control=sequence topic=zod
//ff:what TestApplyMin 테스트
package zod

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestApplyMin(t *testing.T) {
	f := &scanner.Field{Type: "string"}
	ApplyMin(f, ChainMethod{Name: "min", Args: []string{"3"}})
	if f.MinLength == nil || *f.MinLength != 3 {
		t.Fatalf("string min: %v", f.MinLength)
	}
	f2 := &scanner.Field{Type: "number"}
	ApplyMin(f2, ChainMethod{Name: "min", Args: []string{"1"}})
	if f2.Minimum == nil || *f2.Minimum != 1 {
		t.Fatalf("number min: %v", f2.Minimum)
	}
	f3 := &scanner.Field{Type: "string"}
	ApplyMin(f3, ChainMethod{Name: "min"})
	if f3.MinLength != nil {
		t.Fatal("no args should be ignored")
	}
}

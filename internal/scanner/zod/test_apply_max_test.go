//ff:func feature=scan type=test control=sequence topic=zod
//ff:what TestApplyMax 테스트
package zod

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestApplyMax(t *testing.T) {
	f := &scanner.Field{Type: "string"}
	ApplyMax(f, ChainMethod{Name: "max", Args: []string{"10"}})
	if f.MaxLength == nil || *f.MaxLength != 10 {
		t.Fatalf("got %v", f.MaxLength)
	}
	f2 := &scanner.Field{Type: "number"}
	ApplyMax(f2, ChainMethod{Name: "max", Args: []string{"99"}})
	if f2.Maximum == nil || *f2.Maximum != 99 {
		t.Fatalf("got %v", f2.Maximum)
	}
}

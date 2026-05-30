//ff:func feature=scan type=test control=sequence topic=echo
//ff:what TestBuildEndpointIndex_Round5 테스트
package echo

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestBuildEndpointIndex_Round5(t *testing.T) {
	eps := []scanner.Endpoint{
		{File: "a.go", Line: 1},
		{File: "b.go", Line: 2},
	}
	m := buildEndpointIndex(eps)
	if len(m) != 2 {
		t.Fatalf("expected 2 entries, got %d", len(m))
	}
	if m[struct {
		file string
		line int
	}{"a.go", 1}] != 0 {
		t.Fatalf("index mismatch: %v", m)
	}
}

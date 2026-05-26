//ff:func feature=scan type=test control=sequence
//ff:what buildEndpointIndex 전 분기 테스트
package gogin

import (
	"testing"

	"github.com/park-jun-woo/juicer/internal/scanner"
)

func TestBuildEndpointIndex(t *testing.T) {
	eps := []scanner.Endpoint{
		{File: "a.go", Line: 1},
		{File: "b.go", Line: 2},
	}
	idx := buildEndpointIndex(eps)
	if len(idx) != 2 {
		t.Fatalf("expected 2 entries, got %d", len(idx))
	}

	// empty
	idx2 := buildEndpointIndex(nil)
	if len(idx2) != 0 {
		t.Fatal("expected empty map")
	}
}

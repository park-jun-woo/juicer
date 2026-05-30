//ff:func feature=scan type=test control=sequence
//ff:what buildEndpointIndex — (file,line) 인덱스 구축 테스트
package fiber

import (
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestBuildEndpointIndex(t *testing.T) {
	eps := []scanner.Endpoint{
		{File: "a.go", Line: 10},
		{File: "b.go", Line: 20},
	}
	m := buildEndpointIndex(eps)
	if m[struct{ file string; line int }{"a.go", 10}] != 0 {
		t.Errorf("a.go:10 should map to 0")
	}
	if m[struct{ file string; line int }{"b.go", 20}] != 1 {
		t.Errorf("b.go:20 should map to 1")
	}
	if len(m) != 2 {
		t.Errorf("expected 2 entries, got %d", len(m))
	}
}

func TestBuildEndpointIndex_Empty(t *testing.T) {
	if len(buildEndpointIndex(nil)) != 0 {
		t.Error("expected empty map")
	}
}

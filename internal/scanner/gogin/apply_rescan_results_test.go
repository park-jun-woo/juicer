//ff:func feature=scan type=test control=sequence
//ff:what applyRescanResults 전 분기 테스트
package gogin

import (
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestApplyRescanResults(t *testing.T) {
	type key = struct {
		file string
		line int
	}

	eps := []scanner.Endpoint{
		{File: "main.go", Line: 10, Path: "/api/v1/users"},
	}

	ctx := &groupArgCtx{
		endpoints: []scanner.Endpoint{
			{File: "main.go", Line: 10, Path: "/old"},
		},
		epIndex: map[key]int{
			{file: "main.go", line: 10}: 0,
		},
	}

	applyRescanResults(eps, ctx)
	if ctx.endpoints[0].Path != "/api/v1/users" {
		t.Fatalf("expected updated path, got %q", ctx.endpoints[0].Path)
	}

	// not found in index -> skip
	eps2 := []scanner.Endpoint{
		{File: "other.go", Line: 99, Path: "/nope"},
	}
	applyRescanResults(eps2, ctx)

	// with middleware
	eps3 := []scanner.Endpoint{
		{File: "main.go", Line: 10, Path: "/mw", Middleware: []string{"auth"}},
	}
	applyRescanResults(eps3, ctx)
	if len(ctx.endpoints[0].Middleware) != 1 {
		t.Fatal("expected middleware to be applied")
	}
}

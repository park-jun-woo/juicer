//ff:func feature=scan type=test control=sequence
//ff:what applyRescanResults — 재스캔 결과 반영 테스트
package fiber

import (
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestApplyRescanResults(t *testing.T) {
	ctx := &groupArgCtx{
		endpoints: []scanner.Endpoint{
			{File: "a.go", Line: 10, Path: "/old"},
			{File: "b.go", Line: 20, Path: "/keep"},
		},
		epIndex: map[struct {
			file string
			line int
		}]int{
			{"a.go", 10}: 0,
			{"b.go", 20}: 1,
		},
	}

	rescanned := []scanner.Endpoint{
		// matched with middleware -> Path + Middleware updated
		{File: "a.go", Line: 10, Path: "/new", Middleware: []string{"auth"}},
		// matched without middleware -> only Path updated
		{File: "b.go", Line: 20, Path: "/updated"},
		// unmatched key -> skipped
		{File: "z.go", Line: 99, Path: "/ignored"},
	}

	applyRescanResults(rescanned, ctx)

	if ctx.endpoints[0].Path != "/new" || len(ctx.endpoints[0].Middleware) != 1 {
		t.Errorf("endpoint 0 = %+v", ctx.endpoints[0])
	}
	if ctx.endpoints[1].Path != "/updated" || len(ctx.endpoints[1].Middleware) != 0 {
		t.Errorf("endpoint 1 = %+v", ctx.endpoints[1])
	}
}

//ff:func feature=scan type=test control=sequence topic=hono
//ff:what buildRequest 테스트
package hono

import (
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner/zod"
)

func TestBuildRequest_PathParams(t *testing.T) {
	r := routeInfo{Method: "GET"}
	ctx := minimalCtx()
	fi := &fileInfo{Path: "/abs/x.ts"}
	req := buildRequest(r, []string{"id", "slug"}, ctx, fi)
	if req == nil || len(req.PathParams) != 2 {
		t.Fatalf("expected 2 path params, got %+v", req)
	}
}

func TestBuildRequest_NoContent(t *testing.T) {
	r := routeInfo{Method: "GET"}
	ctx := minimalCtx()
	fi := &fileInfo{Path: "/abs/x.ts"}
	if req := buildRequest(r, nil, ctx, fi); req != nil {
		t.Fatalf("expected nil for no content, got %+v", req)
	}
}

func TestBuildRequest_ZodValidator(t *testing.T) {
	fi := mustParse(t, []byte(`const s = z.object({ name: z.string() });`+"\n"))
	calls := findAllByType(fi.Root, "call_expression")
	if len(calls) == 0 {
		t.Skip("no call_expression")
	}
	// the outer z.object(...) call is the first call_expression
	r := routeInfo{
		Method: "POST",
		ZodValidators: []zod.ValidatorInfo{
			{Target: "json", SchemaNode: calls[0]},
		},
	}
	ctx := minimalCtx()
	req := buildRequest(r, nil, ctx, fi)
	if req == nil || req.Body == nil || len(req.Body.Fields) == 0 {
		t.Fatalf("expected body from zod validator, got %+v", req)
	}
}

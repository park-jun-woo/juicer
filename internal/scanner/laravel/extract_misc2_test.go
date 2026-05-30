//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what applyControllerParamTypes / extractGroupModifier / extractChainedGroupIfPresent / extractOneResponse 테스트
package laravel

import (
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestApplyControllerParamTypes_NilCM(t *testing.T) {
	pp := []scanner.Param{{Name: "id", Type: "string"}}
	got := applyControllerParamTypes(pp, nil)
	if len(got) != 1 || got[0].Type != "string" {
		t.Fatalf("got %+v", got)
	}
}

func TestApplyControllerParamTypes_Typed(t *testing.T) {
	cm := &controllerMethod{params: []methodParam{{name: "id", typeName: "int"}}}
	pp := []scanner.Param{{Name: "id", Type: "string"}, {Name: "slug", Type: "string"}}
	got := applyControllerParamTypes(pp, cm)
	if got[0].Type != "integer" {
		t.Fatalf("id type: %+v", got[0])
	}
	if got[1].Type != "string" {
		t.Fatalf("slug unchanged: %+v", got[1])
	}
}

func TestExtractGroupModifier_Prefix(t *testing.T) {
	fi := mustParsePHP(t, `<?php Route::prefix('admin');`)
	call := findAllByType(fi.root, "scoped_call_expression")[0]
	p, mw := extractGroupModifier(call, fi)
	if p != "admin" || mw != nil {
		t.Fatalf("got %q %v", p, mw)
	}
}

func TestExtractGroupModifier_Middleware(t *testing.T) {
	fi := mustParsePHP(t, `<?php Route::middleware(['auth']);`)
	call := findAllByType(fi.root, "scoped_call_expression")[0]
	p, mw := extractGroupModifier(call, fi)
	if p != "" || len(mw) != 1 {
		t.Fatalf("got %q %v", p, mw)
	}
}

func TestExtractGroupModifier_Other(t *testing.T) {
	fi := mustParsePHP(t, `<?php Route::name('admin');`)
	call := findAllByType(fi.root, "scoped_call_expression")[0]
	p, mw := extractGroupModifier(call, fi)
	if p != "" || mw != nil {
		t.Fatalf("got %q %v", p, mw)
	}
}

func TestExtractChainedGroupIfPresent_NoInner(t *testing.T) {
	fi := mustParsePHP(t, `<?php Route::get('/x', [C::class, 'm']);`)
	// scoped call has no inner member_call_expression
	scoped := findAllByType(fi.root, "scoped_call_expression")[0]
	if r := extractChainedGroupIfPresent(scoped, fi, "", nil); r != nil {
		t.Fatalf("expected nil, got %+v", r)
	}
}

func TestExtractOneResponse_JSON(t *testing.T) {
	fi := mustParsePHP(t, `<?php class C { public function m() { return response()->json(['ok' => true], 201); } }`)
	rets := findAllByType(fi.root, "return_statement")
	if len(rets) == 0 {
		t.Fatal("no return")
	}
	resp := extractOneResponse(t.TempDir(), rets[0], fi.src, map[string]*fileInfo{})
	if resp == nil || resp.Status != "201" {
		t.Fatalf("got %+v", resp)
	}
}

func TestExtractOneResponse_NoContent(t *testing.T) {
	fi := mustParsePHP(t, `<?php class C { public function m() { return response()->noContent(); } }`)
	rets := findAllByType(fi.root, "return_statement")
	resp := extractOneResponse(t.TempDir(), rets[0], fi.src, map[string]*fileInfo{})
	if resp == nil || resp.Status != "204" {
		t.Fatalf("got %+v", resp)
	}
}

func TestExtractOneResponse_Plain(t *testing.T) {
	fi := mustParsePHP(t, `<?php class C { public function m() { return $x; } }`)
	rets := findAllByType(fi.root, "return_statement")
	if resp := extractOneResponse(t.TempDir(), rets[0], fi.src, map[string]*fileInfo{}); resp != nil {
		t.Fatalf("expected nil for plain return, got %+v", resp)
	}
}

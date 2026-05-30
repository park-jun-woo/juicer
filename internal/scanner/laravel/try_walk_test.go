//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what tryJSONResponse / tryNoContentResponse / tryResourceResponse / walkChain / walkMemberChain / walkNodes 테스트
package laravel

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
)

func TestTryJSONResponse(t *testing.T) {
	fi := mustParsePHP(t, `<?php class C { public function m() { return response()->json([], 201); } }`)
	ret := findAllByType(fi.root, "return_statement")[0]
	resp := tryJSONResponse(ret, fi.src, nodeText(ret, fi.src))
	if resp == nil || resp.Status != "201" || resp.Kind != "json" {
		t.Fatalf("got %+v", resp)
	}
}

func TestTryJSONResponse_NotJSON(t *testing.T) {
	if resp := tryJSONResponse(nil, nil, "return $x;"); resp != nil {
		t.Fatalf("expected nil, got %+v", resp)
	}
}

func TestTryNoContentResponse(t *testing.T) {
	resp := tryNoContentResponse("return response()->noContent();")
	if resp == nil || resp.Status != "204" || resp.Kind != "empty" {
		t.Fatalf("got %+v", resp)
	}
}

func TestTryNoContentResponse_Not(t *testing.T) {
	if resp := tryNoContentResponse("return $x;"); resp != nil {
		t.Fatalf("expected nil, got %+v", resp)
	}
}

func TestTryResourceResponse(t *testing.T) {
	fi := mustParsePHP(t, `<?php class C { public function show() { return new UserResource($u); } }`)
	ret := findAllByType(fi.root, "return_statement")[0]
	resp := tryResourceResponse(t.TempDir(), ret, fi.src, map[string]*fileInfo{})
	if resp == nil || resp.TypeName != "UserResource" {
		t.Fatalf("got %+v", resp)
	}
}

func TestTryResourceResponse_None(t *testing.T) {
	fi := mustParsePHP(t, `<?php class C { public function show() { return $x; } }`)
	ret := findAllByType(fi.root, "return_statement")[0]
	if resp := tryResourceResponse(t.TempDir(), ret, fi.src, map[string]*fileInfo{}); resp != nil {
		t.Fatalf("expected nil, got %+v", resp)
	}
}

func TestWalkChain_Scoped(t *testing.T) {
	fi := mustParsePHP(t, `<?php Route::prefix('admin');`)
	call := findAllByType(fi.root, "scoped_call_expression")[0]
	prefix := ""
	var mw []string
	walkChain(call, fi, &prefix, &mw)
	if prefix != "admin" {
		t.Fatalf("prefix: %q", prefix)
	}
}

func TestWalkChain_Member(t *testing.T) {
	fi := mustParsePHP(t, `<?php Route::prefix('admin')->middleware('auth');`)
	mcs := findAllByType(fi.root, "member_call_expression")
	if len(mcs) == 0 {
		t.Skip("no member call")
	}
	prefix := ""
	var mw []string
	walkChain(mcs[0], fi, &prefix, &mw)
	if prefix != "admin" || len(mw) == 0 {
		t.Fatalf("prefix=%q mw=%v", prefix, mw)
	}
}

func TestWalkChain_Other(t *testing.T) {
	fi := mustParsePHP(t, `<?php $x = 1;`)
	ints := findAllByType(fi.root, "integer")
	if len(ints) == 0 {
		t.Skip("no int")
	}
	prefix := ""
	var mw []string
	walkChain(ints[0], fi, &prefix, &mw)
	if prefix != "" || mw != nil {
		t.Fatalf("unexpected: %q %v", prefix, mw)
	}
}

func TestWalkNodes(t *testing.T) {
	fi := mustParsePHP(t, `<?php foo(bar());`)
	count := 0
	walkNodes(fi.root, func(n *sitter.Node) {
		if n.Type() == "function_call_expression" {
			count++
		}
	})
	if count != 2 {
		t.Fatalf("expected 2 calls, got %d", count)
	}
}

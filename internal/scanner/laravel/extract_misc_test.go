//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what expandAPIResource / extractClassReference / extractControllerAction / extractMethodParams / extractMiddlewareArgs 테스트
package laravel

import "testing"

func TestExpandAPIResource(t *testing.T) {
	fi := mustParsePHP(t, `<?php Route::apiResource('posts', PostController::class);`)
	calls := findAllByType(fi.root, "scoped_call_expression")
	if len(calls) == 0 {
		t.Fatal("no call")
	}
	routes := expandAPIResource(calls[0], fi, "api", nil)
	if len(routes) == 0 {
		t.Fatal("expected CRUD routes")
	}
	for _, r := range routes {
		if r.controller != "PostController" {
			t.Errorf("controller %q", r.controller)
		}
	}
}

func TestExpandAPIResource_NotApiResource(t *testing.T) {
	fi := mustParsePHP(t, `<?php Route::get('/x', [C::class, 'm']);`)
	calls := findAllByType(fi.root, "scoped_call_expression")
	if r := expandAPIResource(calls[0], fi, "", nil); r != nil {
		t.Fatalf("expected nil, got %+v", r)
	}
}

func TestExtractClassReference(t *testing.T) {
	fi := mustParsePHP(t, `<?php $x = PostController::class;`)
	exprs := findAllByType(fi.root, "class_constant_access_expression")
	if len(exprs) == 0 {
		t.Fatal("no class const access")
	}
	if got := extractClassReference(exprs[0], fi.src); got != "PostController" {
		t.Fatalf("got %q", got)
	}
}

func TestExtractClassReference_NoAccess(t *testing.T) {
	fi := mustParsePHP(t, `<?php $x = 'plain';`)
	str := findAllByType(fi.root, "string")[0]
	if got := extractClassReference(str, fi.src); got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}

func TestExtractControllerAction(t *testing.T) {
	fi := mustParsePHP(t, `<?php $x = [UserController::class, 'show'];`)
	arrs := findAllByType(fi.root, "array_creation_expression")
	ctrl, action := extractControllerAction(arrs[0], fi.src)
	if ctrl != "UserController" || action != "show" {
		t.Fatalf("got %q %q", ctrl, action)
	}
}

func TestExtractControllerAction_NotArray(t *testing.T) {
	fi := mustParsePHP(t, `<?php $x = 'closure';`)
	str := findAllByType(fi.root, "string")[0]
	ctrl, action := extractControllerAction(str, fi.src)
	if ctrl != "" || action != "" {
		t.Fatalf("got %q %q", ctrl, action)
	}
}

func TestExtractMethodParams(t *testing.T) {
	fi := mustParsePHP(t, `<?php class C { public function m(int $id, string $name) {} }`)
	fp := findAllByType(fi.root, "formal_parameters")[0]
	params := extractMethodParams(fp, fi.src)
	if len(params) != 2 || params[0].name != "id" || params[0].typeName != "int" {
		t.Fatalf("got %+v", params)
	}
}

func TestExtractMiddlewareArgs_Single(t *testing.T) {
	fi := mustParsePHP(t, `<?php Route::middleware('auth');`)
	args := findAllByType(fi.root, "arguments")[0]
	mw := extractMiddlewareArgs(args, fi)
	if len(mw) != 1 || mw[0] != "auth" {
		t.Fatalf("got %v", mw)
	}
}

func TestExtractMiddlewareArgs_Array(t *testing.T) {
	fi := mustParsePHP(t, `<?php Route::middleware(['auth', 'throttle']);`)
	args := findAllByType(fi.root, "arguments")[0]
	mw := extractMiddlewareArgs(args, fi)
	if len(mw) != 2 {
		t.Fatalf("got %v", mw)
	}
}

func TestExtractMiddlewareArgs_None(t *testing.T) {
	fi := mustParsePHP(t, `<?php Route::middleware();`)
	args := findAllByType(fi.root, "arguments")[0]
	if mw := extractMiddlewareArgs(args, fi); mw != nil {
		t.Fatalf("expected nil, got %v", mw)
	}
}

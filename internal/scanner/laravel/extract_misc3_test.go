//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what extractClosureBody / extractControllerMethod / extractFieldsFromRulesMethod / extractOneRuleField / extractOneRoute / extractMemberCallRoutes / extractOneGroup / extractGroupArrayModifier / extractResponsesFromMethod 테스트
package laravel

import "testing"

func TestExtractClosureBody_AnonFunc(t *testing.T) {
	fi := mustParsePHP(t, `<?php Route::group([], function () { $x = 1; });`)
	args := findAllByType(fi.root, "arguments")[0]
	if extractClosureBody(args, fi) == nil {
		t.Fatal("expected closure body")
	}
}

func TestExtractClosureBody_None(t *testing.T) {
	fi := mustParsePHP(t, `<?php Route::group([], $cb);`)
	args := findAllByType(fi.root, "arguments")[0]
	if extractClosureBody(args, fi) != nil {
		t.Fatal("expected nil")
	}
}

func TestExtractControllerMethod(t *testing.T) {
	fi := mustParsePHP(t, `<?php class C { public function show(int $id) { return $id; } }`)
	cm := extractControllerMethod(&fi, "show")
	if cm == nil {
		t.Fatal("nil cm")
	}
	if cm.name != "show" || len(cm.params) != 1 || cm.params[0].name != "id" {
		t.Fatalf("got %+v", cm)
	}
	if len(cm.returnNodes) != 1 {
		t.Fatalf("returns: %+v", cm.returnNodes)
	}
}

func TestExtractControllerMethod_NotFound(t *testing.T) {
	fi := mustParsePHP(t, `<?php class C { public function index() {} }`)
	if cm := extractControllerMethod(&fi, "missing"); cm != nil {
		t.Fatalf("expected nil, got %+v", cm)
	}
}

func TestExtractFieldsFromRulesMethod(t *testing.T) {
	fi := mustParsePHP(t, `<?php class R { public function rules(): array {
		return [ 'name' => 'required|string', 'age' => 'integer' ];
	} }`)
	method := findAllByType(fi.root, "method_declaration")[0]
	fields := extractFieldsFromRulesMethod(method, fi.src)
	if len(fields) != 2 || fields[0].Name != "name" {
		t.Fatalf("got %+v", fields)
	}
}

func TestExtractFieldsFromRulesMethod_NoArray(t *testing.T) {
	fi := mustParsePHP(t, `<?php class R { public function rules() { return null; } }`)
	method := findAllByType(fi.root, "method_declaration")[0]
	if got := extractFieldsFromRulesMethod(method, fi.src); got != nil {
		t.Fatalf("expected nil, got %+v", got)
	}
}

func TestExtractOneRuleField_ArrayRules(t *testing.T) {
	fi := mustParsePHP(t, `<?php $x = ['role' => ['required', 'in:a,b']];`)
	elems := findAllByType(fi.root, "array_element_initializer")
	field := extractOneRuleField(elems[0], fi.src)
	if field == nil || field.Name != "role" {
		t.Fatalf("got %+v", field)
	}
	if len(field.Enum) != 2 {
		t.Fatalf("enum: %+v", field.Enum)
	}
}

func TestExtractOneRoute_Get(t *testing.T) {
	fi := mustParsePHP(t, `<?php Route::get('/users', [UserController::class, 'index']);`)
	call := findAllByType(fi.root, "scoped_call_expression")[0]
	r := extractOneRoute(call, fi, "api", nil)
	if r == nil || r.method != "GET" || r.path != "/api/users" {
		t.Fatalf("got %+v", r)
	}
}

func TestExtractOneRoute_UnknownMethod(t *testing.T) {
	fi := mustParsePHP(t, `<?php Route::macro('/x', $cb);`)
	call := findAllByType(fi.root, "scoped_call_expression")[0]
	if r := extractOneRoute(call, fi, "", nil); r != nil {
		t.Fatalf("expected nil, got %+v", r)
	}
}

func TestExtractOneRoute_NotRouteScope(t *testing.T) {
	fi := mustParsePHP(t, `<?php Foo::get('/x', $h);`)
	call := findAllByType(fi.root, "scoped_call_expression")[0]
	if r := extractOneRoute(call, fi, "", nil); r != nil {
		t.Fatalf("expected nil, got %+v", r)
	}
}

func TestExtractGroupArrayModifier(t *testing.T) {
	fi := mustParsePHP(t, `<?php $x = ['prefix' => 'admin', 'middleware' => ['auth']];`)
	arr := findAllByType(fi.root, "array_creation_expression")[0]
	prefix, mw := extractGroupArrayModifier(arr, fi)
	if prefix == "" || len(mw) != 1 {
		t.Fatalf("got %q %v", prefix, mw)
	}
}

func TestExtractResponsesFromMethod(t *testing.T) {
	fi := mustParsePHP(t, `<?php class C { public function m() {
		return response()->json(['ok' => true], 201);
	} }`)
	cm := extractControllerMethod(&fi, "m")
	resps := extractResponsesFromMethod(t.TempDir(), cm, map[string]*fileInfo{})
	if len(resps) == 0 || resps[0].Status != "201" {
		t.Fatalf("got %+v", resps)
	}
}

func TestExtractMemberCallRoutes_Group(t *testing.T) {
	fi := mustParsePHP(t, `<?php Route::prefix('admin')->group(function () {
		Route::get('/users', [C::class, 'index']);
	});`)
	mcs := findAllByType(fi.root, "member_call_expression")
	if len(mcs) == 0 {
		t.Skip("no member call")
	}
	routes := extractMemberCallRoutes(mcs[0], fi, "", nil)
	if len(routes) == 0 {
		t.Fatalf("expected routes, got %+v", routes)
	}
}

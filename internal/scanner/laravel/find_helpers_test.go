//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what extractRuleStrings / extractStringArray / extractToArrayFields / extractRulesFromFile / extractScopedGroup / findChildByType / findClassConstantAccess / findClassMethod / findAnyClassMethod / findFormRequestParam / findFormRequestFile 테스트
package laravel

import "testing"

func TestExtractRuleStrings_Pipe(t *testing.T) {
	fi := mustParsePHP(t, `<?php $x = ['name' => 'required|string|max:255'];`)
	elem := findAllByType(fi.root, "array_element_initializer")[0]
	rules := extractRuleStrings(elem, fi.src)
	if len(rules) != 3 || rules[0] != "required" || rules[2] != "max:255" {
		t.Fatalf("got %v", rules)
	}
}

func TestExtractRuleStrings_Array(t *testing.T) {
	fi := mustParsePHP(t, `<?php $x = ['role' => ['required', 'in:a,b']];`)
	elem := findAllByType(fi.root, "array_element_initializer")[0]
	rules := extractRuleStrings(elem, fi.src)
	if len(rules) != 2 {
		t.Fatalf("got %v", rules)
	}
}

func TestExtractRuleStrings_NoValue(t *testing.T) {
	fi := mustParsePHP(t, `<?php $x = ['name'];`)
	elem := findAllByType(fi.root, "array_element_initializer")[0]
	if rules := extractRuleStrings(elem, fi.src); rules != nil {
		t.Fatalf("expected nil, got %v", rules)
	}
}

func TestExtractStringArray(t *testing.T) {
	fi := mustParsePHP(t, `<?php $x = ['a', 'b', 'c'];`)
	arr := findAllByType(fi.root, "array_creation_expression")[0]
	got := extractStringArray(arr, fi.src)
	if len(got) != 3 || got[0] != "a" {
		t.Fatalf("got %v", got)
	}
}

func TestExtractToArrayFields(t *testing.T) {
	fi := mustParsePHP(t, `<?php class UserResource {
		public function toArray($request) {
			return [ 'id' => $this->id, 'name' => $this->name ];
		}
	}`)
	fields := extractToArrayFields(&fi, "UserResource")
	if len(fields) != 2 {
		t.Fatalf("got %+v", fields)
	}
}

func TestExtractToArrayFields_NoMethod(t *testing.T) {
	fi := mustParsePHP(t, `<?php class R {}`)
	if got := extractToArrayFields(&fi, "R"); got != nil {
		t.Fatalf("expected nil, got %+v", got)
	}
}

func TestExtractRulesFromFile(t *testing.T) {
	fi := mustParsePHP(t, `<?php class StoreReq {
		public function rules(): array { return [ 'name' => 'required|string' ]; }
	}`)
	fields := extractRulesFromFile(&fi, "StoreReq")
	if len(fields) != 1 || fields[0].Name != "name" {
		t.Fatalf("got %+v", fields)
	}
}

func TestExtractRulesFromFile_NoMethod(t *testing.T) {
	fi := mustParsePHP(t, `<?php class R {}`)
	if got := extractRulesFromFile(&fi, "R"); got != nil {
		t.Fatalf("expected nil, got %+v", got)
	}
}

func TestExtractScopedGroup(t *testing.T) {
	fi := mustParsePHP(t, `<?php Route::group(['prefix' => 'admin'], function () {
		Route::get('/users', [C::class, 'index']);
	});`)
	call := findAllByType(fi.root, "scoped_call_expression")[0]
	routes := extractScopedGroup(call, fi, "", nil)
	if len(routes) == 0 {
		t.Fatalf("expected routes, got %+v", routes)
	}
	if routes[0].path != "/admin/users" {
		t.Fatalf("path: %q", routes[0].path)
	}
}

func TestExtractScopedGroup_NotGroup(t *testing.T) {
	fi := mustParsePHP(t, `<?php Route::get('/x', $h);`)
	call := findAllByType(fi.root, "scoped_call_expression")[0]
	if r := extractScopedGroup(call, fi, "", nil); r != nil {
		t.Fatalf("expected nil, got %+v", r)
	}
}

func TestExtractScopedGroup_NotRouteScope(t *testing.T) {
	fi := mustParsePHP(t, `<?php Foo::group(['prefix'=>'x'], fn() => null);`)
	call := findAllByType(fi.root, "scoped_call_expression")[0]
	if r := extractScopedGroup(call, fi, "", nil); r != nil {
		t.Fatalf("expected nil, got %+v", r)
	}
}

func TestFindChildByType(t *testing.T) {
	fi := mustParsePHP(t, `<?php foo();`)
	calls := findAllByType(fi.root, "function_call_expression")
	if len(calls) == 0 {
		t.Skip("no call")
	}
	if findChildByType(calls[0], "arguments") == nil {
		t.Fatal("expected arguments")
	}
	if findChildByType(calls[0], "object") != nil {
		t.Fatal("expected nil for missing type")
	}
}

func TestFindClassConstantAccess_Direct(t *testing.T) {
	fi := mustParsePHP(t, `<?php $x = Foo::class;`)
	// the assignment value node directly holds class_constant_access
	exprs := findAllByType(fi.root, "class_constant_access_expression")
	if len(exprs) == 0 {
		t.Fatal("no expr")
	}
	parent := exprs[0]
	if got := findClassConstantAccess(parent); got == nil {
		t.Fatal("expected to find self/child access")
	}
}

func TestFindClassConstantAccess_None(t *testing.T) {
	fi := mustParsePHP(t, `<?php $x = 'plain';`)
	str := findAllByType(fi.root, "string")[0]
	if findClassConstantAccess(str) != nil {
		t.Fatal("expected nil")
	}
}

func TestFindClassMethod(t *testing.T) {
	fi := mustParsePHP(t, `<?php class A { public function foo() {} } class B { public function bar() {} }`)
	if findClassMethod(&fi, "B", "bar") == nil {
		t.Fatal("expected to find B::bar")
	}
	if findClassMethod(&fi, "A", "bar") != nil {
		t.Fatal("A has no bar")
	}
	if findClassMethod(&fi, "Missing", "foo") != nil {
		t.Fatal("missing class")
	}
}

func TestFindAnyClassMethod(t *testing.T) {
	fi := mustParsePHP(t, `<?php class A {} class B { public function bar() {} }`)
	if findAnyClassMethod(&fi, "bar") == nil {
		t.Fatal("expected to find bar in any class")
	}
	if findAnyClassMethod(&fi, "nope") != nil {
		t.Fatal("expected nil")
	}
}

func TestFindFormRequestParam(t *testing.T) {
	params := []methodParam{
		{name: "id", typeName: "int"},
		{name: "request", typeName: "StoreUserRequest"},
	}
	if got := findFormRequestParam(params); got != "StoreUserRequest" {
		t.Fatalf("got %q", got)
	}
	if got := findFormRequestParam([]methodParam{{name: "id", typeName: "int"}}); got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}

func TestFindFormRequestFile_Parsed(t *testing.T) {
	fi := mustParsePHP(t, `<?php class StoreReq {}`)
	parsed := map[string]*fileInfo{"x.php": &fi}
	if findFormRequestFile("/root", "StoreReq", parsed) == nil {
		t.Fatal("expected to find via parsed files")
	}
}

func TestFindFormRequestFile_PSR4(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "app/Http/Requests/StoreReq.php", `<?php class StoreReq {}`)
	if findFormRequestFile(dir, "StoreReq", map[string]*fileInfo{}) == nil {
		t.Fatal("expected to find via PSR-4 path")
	}
}

func TestFindFormRequestFile_NotFound(t *testing.T) {
	dir := t.TempDir()
	if findFormRequestFile(dir, "Missing", map[string]*fileInfo{}) != nil {
		t.Fatal("expected nil")
	}
}

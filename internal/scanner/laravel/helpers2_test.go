//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what 다수 헬퍼 함수 테스트 (join/merge/node/param/find/json/middleware 등)
package laravel

import (
	"path/filepath"
	"testing"
)

func TestJoinGroupPrefix(t *testing.T) {
	if joinGroupPrefix("", "v1") != "v1" {
		t.Fatal("empty outer")
	}
	if joinGroupPrefix("api", "") != "api" {
		t.Fatal("empty inner")
	}
	if joinGroupPrefix("api", "v1") != "api/v1" {
		t.Fatal("both")
	}
}

func TestIsFormRequestType(t *testing.T) {
	if isFormRequestType("StoreUserRequest") != true {
		t.Fatal("expected true")
	}
	for _, n := range []string{"", "Request", "int", "string", "float", "bool", "array", "User"} {
		if isFormRequestType(n) {
			t.Errorf("expected false for %q", n)
		}
	}
}

func TestLastSegmentSingular(t *testing.T) {
	if got := lastSegmentSingular("users"); got != "user" {
		t.Fatalf("got %q", got)
	}
	if got := lastSegmentSingular("{product}/reviews"); got != "review" {
		t.Fatalf("got %q", got)
	}
}

func TestMergeMiddleware(t *testing.T) {
	if got := mergeMiddleware([]string{"a"}, nil); len(got) != 1 {
		t.Fatal("empty b")
	}
	got := mergeMiddleware([]string{"a", "b"}, []string{"b", "c"})
	if len(got) != 3 {
		t.Fatalf("dedup failed: %v", got)
	}
}

func TestNodeText(t *testing.T) {
	fi := mustParsePHP(t, `<?php $hello;`)
	vars := findAllByType(fi.root, "variable_name")
	if len(vars) == 0 {
		t.Skip("no var")
	}
	if got := nodeText(vars[0], fi.src); got != "$hello" {
		t.Fatalf("got %q", got)
	}
}

func TestParamTypeName(t *testing.T) {
	fi := mustParsePHP(t, `<?php class C { public function m(int $id, FooRequest $r, $untyped) {} }`)
	sps := findAllByType(fi.root, "simple_parameter")
	if len(sps) < 3 {
		t.Fatalf("expected 3 params, got %d", len(sps))
	}
	if got := paramTypeName(sps[0], fi.src); got != "int" {
		t.Fatalf("primitive: %q", got)
	}
	if got := paramTypeName(sps[1], fi.src); got != "FooRequest" {
		t.Fatalf("named: %q", got)
	}
	if got := paramTypeName(sps[2], fi.src); got != "" {
		t.Fatalf("untyped: %q", got)
	}
}

func TestFirstArgString(t *testing.T) {
	fi := mustParsePHP(t, `<?php Route::prefix('admin');`)
	args := findAllByType(fi.root, "arguments")[0]
	s, ok := firstArgString(args, fi.src)
	if !ok || s != "admin" {
		t.Fatalf("got %q %v", s, ok)
	}
}

func TestFirstArgString_NoArgs(t *testing.T) {
	fi := mustParsePHP(t, `<?php foo();`)
	args := findAllByType(fi.root, "arguments")[0]
	if _, ok := firstArgString(args, fi.src); ok {
		t.Fatal("expected false")
	}
}

func TestGroupArrayValue(t *testing.T) {
	fi := mustParsePHP(t, `<?php $x = ['prefix' => 'v1'];`)
	elem := findAllByType(fi.root, "array_element_initializer")[0]
	v := groupArrayValue(elem)
	if v == nil || extractStringContent(v, fi.src) != "v1" {
		t.Fatalf("got %v", v)
	}
}

func TestGroupArrayValue_NoArrow(t *testing.T) {
	fi := mustParsePHP(t, `<?php $x = ['v1'];`)
	elem := findAllByType(fi.root, "array_element_initializer")[0]
	if v := groupArrayValue(elem); v != nil {
		t.Fatalf("expected nil, got %v", v)
	}
}

func TestMemberCallName(t *testing.T) {
	fi := mustParsePHP(t, `<?php $obj->doThing();`)
	mcs := findAllByType(fi.root, "member_call_expression")
	if len(mcs) == 0 {
		t.Skip("no member call")
	}
	if got := memberCallName(mcs[0], fi.src); got != "doThing" {
		t.Fatalf("got %q", got)
	}
}

func TestLastMemberCallName(t *testing.T) {
	fi := mustParsePHP(t, `<?php Route::get('/x')->name('foo');`)
	mcs := findAllByType(fi.root, "member_call_expression")
	if got := lastMemberCallName(mcs[0], fi.src); got != "name" {
		t.Fatalf("got %q", got)
	}
}

func TestJSONResponseStatus(t *testing.T) {
	fi := mustParsePHP(t, `<?php class C { public function m() { return response()->json([], 201); } }`)
	ret := findAllByType(fi.root, "return_statement")[0]
	if got := jsonResponseStatus(ret, fi.src); got != "201" {
		t.Fatalf("got %q", got)
	}
}

func TestJSONResponseStatus_Default(t *testing.T) {
	fi := mustParsePHP(t, `<?php class C { public function m() { return response()->json([]); } }`)
	ret := findAllByType(fi.root, "return_statement")[0]
	if got := jsonResponseStatus(ret, fi.src); got != "200" {
		t.Fatalf("got %q", got)
	}
}

func TestJSONCallStatusCode_WithStatus(t *testing.T) {
	fi := mustParsePHP(t, `<?php $x = response()->json([], 201);`)
	mcs := findAllByType(fi.root, "member_call_expression")
	if got := jsonCallStatusCode(mcs[0], fi.src); got != "201" {
		t.Fatalf("got %q", got)
	}
}

func TestJSONCallStatusCode_OneArg(t *testing.T) {
	fi := mustParsePHP(t, `<?php $x = response()->json([]);`)
	mcs := findAllByType(fi.root, "member_call_expression")
	if got := jsonCallStatusCode(mcs[0], fi.src); got != "" {
		t.Fatalf("expected empty for single arg, got %q", got)
	}
}

func TestJSONCallStatusCode_NotJson(t *testing.T) {
	fi := mustParsePHP(t, `<?php $x = $obj->other(1, 2);`)
	mcs := findAllByType(fi.root, "member_call_expression")
	if got := jsonCallStatusCode(mcs[0], fi.src); got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}

func TestMiddlewareValues_String(t *testing.T) {
	fi := mustParsePHP(t, `<?php $x = 'auth';`)
	str := findAllByType(fi.root, "string")[0]
	if got := middlewareValues(str, fi); len(got) != 1 || got[0] != "auth" {
		t.Fatalf("got %v", got)
	}
}

func TestMiddlewareValues_Array(t *testing.T) {
	fi := mustParsePHP(t, `<?php $x = ['auth', 'throttle'];`)
	arr := findAllByType(fi.root, "array_creation_expression")[0]
	if got := middlewareValues(arr, fi); len(got) != 2 {
		t.Fatalf("got %v", got)
	}
}

func TestMethodReturnArray(t *testing.T) {
	fi := mustParsePHP(t, `<?php class C { public function m() { return ['a' => 1]; } }`)
	method := findAllByType(fi.root, "method_declaration")[0]
	if methodReturnArray(method) == nil {
		t.Fatal("expected array")
	}
}

func TestMethodReturnArray_NoReturn(t *testing.T) {
	fi := mustParsePHP(t, `<?php class C { public function m() { $x = 1; } }`)
	method := findAllByType(fi.root, "method_declaration")[0]
	if methodReturnArray(method) != nil {
		t.Fatal("expected nil")
	}
}

func TestFindMethodInClass(t *testing.T) {
	fi := mustParsePHP(t, `<?php class C { public function foo() {} public function bar() {} }`)
	cls := findAllByType(fi.root, "class_declaration")[0]
	if findMethodInClass(cls, fi.src, "bar") == nil {
		t.Fatal("expected bar")
	}
	if findMethodInClass(cls, fi.src, "nope") != nil {
		t.Fatal("expected nil")
	}
}

func TestGroupClosureBody(t *testing.T) {
	fi := mustParsePHP(t, `<?php Route::prefix('x')->group(function () { $a = 1; });`)
	mcs := findAllByType(fi.root, "member_call_expression")
	if len(mcs) == 0 {
		t.Skip("no member call")
	}
	if groupClosureBody(mcs[0], fi) == nil {
		t.Fatal("expected closure body")
	}
}

func TestIsGroupCallArgument(t *testing.T) {
	fi := mustParsePHP(t, `<?php Route::prefix('x')->group(function () { Route::get('/y', $h); });`)
	closures := findAllByType(fi.root, "anonymous_function_creation_expression")
	if len(closures) == 0 {
		t.Skip("no closure")
	}
	if !isGroupCallArgument(closures[0], fi) {
		t.Fatal("expected true for group closure")
	}
}

func TestFindResourceFile_Parsed(t *testing.T) {
	fi := mustParsePHP(t, `<?php class UserResource {}`)
	parsed := map[string]*fileInfo{"x.php": &fi}
	if findResourceFile("/root", "UserResource", parsed) == nil {
		t.Fatal("expected found")
	}
}

func TestFindResourceFile_NotFound(t *testing.T) {
	if findResourceFile(t.TempDir(), "Missing", map[string]*fileInfo{}) != nil {
		t.Fatal("expected nil")
	}
}

func TestFindPHPFiles(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "app/Foo.php", "<?php")
	writeFile(t, dir, "vendor/lib.php", "<?php")
	writeFile(t, dir, "Foo.test.php", "<?php")
	files, err := findPHPFiles(dir)
	if err != nil {
		t.Fatal(err)
	}
	for _, f := range files {
		rel, _ := filepath.Rel(dir, f)
		if rel != filepath.Join("app", "Foo.php") {
			t.Errorf("unexpected file: %s", rel)
		}
	}
}

func TestLaravelRulesToField(t *testing.T) {
	f := laravelRulesToField("email", []string{"required", "email"})
	if f.Name != "email" || f.Type != "string" {
		t.Fatalf("got %+v", f)
	}
	f2 := laravelRulesToField("age", []string{"integer", "min:0"})
	if f2.Type != "integer" {
		t.Fatalf("got %+v", f2)
	}
}

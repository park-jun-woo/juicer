//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what parseAllPHPFiles / parseControllerCandidate / parseSimpleParam / resolveArrayNode / resolveController / resolveRouteController / resourceTypeName / sameNode / secondScopedName / tryCollectionResponse 테스트
package laravel

import (
	"path/filepath"
	"testing"
)

func TestParseAllPHPFiles(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "a.php", "<?php class A {}")
	writeFile(t, dir, "b.php", "<?php class B {}")
	parsed := parseAllPHPFiles(dir, []string{
		filepath.Join(dir, "a.php"),
		filepath.Join(dir, "b.php"),
		filepath.Join(dir, "missing.php"), // skipped on parse error
	})
	if len(parsed) != 2 {
		t.Fatalf("expected 2, got %d", len(parsed))
	}
}

func TestParseControllerCandidate(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "C.php", "<?php class C {}")
	if parseControllerCandidate(dir, filepath.Join(dir, "C.php")) == nil {
		t.Fatal("expected fileInfo")
	}
	if parseControllerCandidate(dir, filepath.Join(dir, "missing.php")) != nil {
		t.Fatal("expected nil for missing")
	}
}

func TestParseSimpleParam(t *testing.T) {
	fi := mustParsePHP(t, `<?php class C { public function m(int $id) {} }`)
	sp := findAllByType(fi.root, "simple_parameter")[0]
	mp := parseSimpleParam(sp, fi.src)
	if mp.name != "id" || mp.typeName != "int" {
		t.Fatalf("got %+v", mp)
	}
}

func TestResolveArrayNode(t *testing.T) {
	// node containing an array
	fi := mustParsePHP(t, `<?php $x = [1, 2];`)
	arrs := findAllByType(fi.root, "array_creation_expression")
	if resolveArrayNode(arrs[0]) == nil {
		t.Fatal("self array node")
	}
	// non-array node
	str := findAllByType(fi.root, "integer")
	if len(str) > 0 && resolveArrayNode(str[0]) != nil {
		t.Fatal("expected nil for non-array")
	}
}

func TestResolveController_Parsed(t *testing.T) {
	fi := mustParsePHP(t, `<?php class UserController {}`)
	parsed := map[string]*fileInfo{"x.php": &fi}
	if resolveController("/root", "UserController", parsed) == nil {
		t.Fatal("expected via parsed")
	}
}

func TestResolveController_Path(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "app/Http/Controllers/Api/PostController.php", `<?php class PostController {}`)
	if resolveController(dir, "PostController", map[string]*fileInfo{}) == nil {
		t.Fatal("expected via Api path")
	}
}

func TestResolveController_NotFound(t *testing.T) {
	if resolveController(t.TempDir(), "Missing", map[string]*fileInfo{}) != nil {
		t.Fatal("expected nil")
	}
}

func TestResolveRouteController(t *testing.T) {
	fi := mustParsePHP(t, `<?php class UserController { public function show(int $id) { return $id; } }`)
	parsed := map[string]*fileInfo{"x.php": &fi}
	ri := routeInfo{controller: "UserController", action: "show"}
	cm := resolveRouteController("/root", ri, parsed)
	if cm == nil || cm.name != "show" {
		t.Fatalf("got %+v", cm)
	}
}

func TestResolveRouteController_NoControllerOrAction(t *testing.T) {
	if resolveRouteController("/root", routeInfo{}, map[string]*fileInfo{}) != nil {
		t.Fatal("expected nil for empty controller/action")
	}
}

func TestResolveRouteController_Unresolved(t *testing.T) {
	ri := routeInfo{controller: "Missing", action: "show"}
	if resolveRouteController(t.TempDir(), ri, map[string]*fileInfo{}) != nil {
		t.Fatal("expected nil when controller not found")
	}
}

func TestResourceTypeName(t *testing.T) {
	fi := mustParsePHP(t, `<?php $x = new UserResource($u);`)
	ocs := findAllByType(fi.root, "object_creation_expression")
	if len(ocs) == 0 {
		t.Skip("no object creation")
	}
	if got := resourceTypeName(ocs[0], fi.src); got != "UserResource" {
		t.Fatalf("got %q", got)
	}
}

func TestResourceTypeName_NotResource(t *testing.T) {
	fi := mustParsePHP(t, `<?php $x = new Foo($u);`)
	ocs := findAllByType(fi.root, "object_creation_expression")
	if len(ocs) == 0 {
		t.Skip("no object creation")
	}
	if got := resourceTypeName(ocs[0], fi.src); got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}

func TestSameNode(t *testing.T) {
	fi := mustParsePHP(t, `<?php foo();`)
	calls := findAllByType(fi.root, "function_call_expression")
	if len(calls) == 0 {
		t.Skip("no call")
	}
	if !sameNode(calls[0], calls[0]) {
		t.Fatal("node should equal itself")
	}
	names := findAllByType(fi.root, "name")
	if len(names) > 0 && sameNode(calls[0], names[0]) {
		t.Fatal("different nodes should not be equal")
	}
}

func TestSecondScopedName(t *testing.T) {
	fi := mustParsePHP(t, `<?php Route::get('/x', $h);`)
	call := findAllByType(fi.root, "scoped_call_expression")[0]
	if got := secondScopedName(call, fi.src); got != "get" {
		t.Fatalf("got %q", got)
	}
}

func TestTryCollectionResponse(t *testing.T) {
	fi := mustParsePHP(t, `<?php class C { public function index() {
		return UserResource::collection($users);
	} }`)
	ret := findAllByType(fi.root, "return_statement")[0]
	resp := tryCollectionResponse(t.TempDir(), ret, fi.src, map[string]*fileInfo{})
	if resp == nil || resp.TypeName != "[]UserResource" {
		t.Fatalf("got %+v", resp)
	}
}

func TestTryCollectionResponse_None(t *testing.T) {
	fi := mustParsePHP(t, `<?php class C { public function index() { return $x; } }`)
	ret := findAllByType(fi.root, "return_statement")[0]
	if resp := tryCollectionResponse(t.TempDir(), ret, fi.src, map[string]*fileInfo{}); resp != nil {
		t.Fatalf("expected nil, got %+v", resp)
	}
}

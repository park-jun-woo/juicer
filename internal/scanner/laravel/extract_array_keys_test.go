//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what extractArrayKeys / controllerClassName / collectAllRoutes 테스트
package laravel

import (
	"path/filepath"
	"testing"
)

func TestExtractArrayKeys(t *testing.T) {
	fi := mustParsePHP(t, `<?php
class R {
    public function toArray($request) {
        return [
            'id' => $this->id,
            'name' => $this->name,
        ];
    }
}`)
	methods := findAllByType(fi.root, "method_declaration")
	if len(methods) == 0 {
		t.Fatal("no method")
	}
	fields := extractArrayKeys(methods[0], fi.src)
	if len(fields) != 2 || fields[0].Name != "id" || fields[1].Name != "name" {
		t.Fatalf("got %+v", fields)
	}
	if fields[0].Type != "string" {
		t.Fatalf("default type: %+v", fields[0])
	}
}

func TestExtractArrayKeys_NoReturnArray(t *testing.T) {
	fi := mustParsePHP(t, `<?php class R { public function toArray($r) { return null; } }`)
	methods := findAllByType(fi.root, "method_declaration")
	if got := extractArrayKeys(methods[0], fi.src); got != nil {
		t.Fatalf("expected nil, got %+v", got)
	}
}

func TestControllerClassName(t *testing.T) {
	fi := mustParsePHP(t, `<?php $x = [UserController::class, 'index'];`)
	elems := findAllByType(fi.root, "array_creation_expression")
	if len(elems) == 0 {
		t.Fatal("no array")
	}
	if got := controllerClassName(elems[0], fi.src); got != "UserController" {
		t.Fatalf("got %q", got)
	}
}

func TestControllerClassName_NoClassAccess(t *testing.T) {
	fi := mustParsePHP(t, `<?php $x = ['a', 'b'];`)
	elems := findAllByType(fi.root, "array_creation_expression")
	if got := controllerClassName(elems[0], fi.src); got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}

func TestCollectAllRoutes(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "routes/api.php", `<?php
Route::get('/users', [UserController::class, 'index']);
Route::post('/users', [UserController::class, 'store']);
`)
	fi, err := parseFile(dir, filepath.Join(dir, "routes/api.php"))
	if err != nil {
		t.Fatal(err)
	}
	parsed := map[string]*fileInfo{"routes/api.php": fi}
	routes := collectAllRoutes(parsed)
	if len(routes) != 2 {
		t.Fatalf("expected 2 routes, got %d", len(routes))
	}
	for _, r := range routes {
		if r.path != "/api/users" {
			t.Errorf("expected api prefix, got %q", r.path)
		}
	}
}

func TestCollectAllRoutes_NoRouteFiles(t *testing.T) {
	if routes := collectAllRoutes(map[string]*fileInfo{}); len(routes) != 0 {
		t.Fatalf("expected none, got %d", len(routes))
	}
}

func TestCollectAllRoutes_WebPrefix(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "routes/web.php", `<?php
Route::get('/home', [HomeController::class, 'index']);
`)
	fi, _ := parseFile(dir, filepath.Join(dir, "routes/web.php"))
	parsed := map[string]*fileInfo{"routes/web.php": fi}
	routes := collectAllRoutes(parsed)
	if len(routes) != 1 || routes[0].path != "/home" {
		t.Fatalf("web route: %+v", routes)
	}
}

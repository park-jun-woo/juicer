//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what buildEndpoints / buildOneEndpoint / buildRequest 테스트
package laravel

import (
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestBuildOneEndpoint_Closure(t *testing.T) {
	ri := routeInfo{method: "GET", path: "/health", file: "routes/api.php", line: 5, middleware: []string{"auth"}}
	ep := buildOneEndpoint(t.TempDir(), ri, map[string]*fileInfo{})
	if ep.Method != "GET" || ep.Path != "/health" || ep.Handler != "closure" {
		t.Fatalf("got %+v", ep)
	}
	if ep.Line != 5 || len(ep.Middleware) != 1 {
		t.Fatalf("meta: %+v", ep)
	}
}

func TestBuildOneEndpoint_WithController(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "app/Http/Controllers/UserController.php", `<?php
namespace App\Http\Controllers;
use App\Http\Requests\StoreUserRequest;
class UserController extends Controller {
    public function store(StoreUserRequest $request) {
        return response()->json(['ok' => true], 201);
    }
    public function show(int $id) {
        return response()->json(['id' => $id]);
    }
}
`)
	writeFile(t, dir, "app/Http/Requests/StoreUserRequest.php", `<?php
namespace App\Http\Requests;
use Illuminate\Foundation\Http\FormRequest;
class StoreUserRequest extends FormRequest {
    public function rules(): array { return [ 'name' => 'required|string' ]; }
}
`)
	parsed := map[string]*fileInfo{}

	store := routeInfo{method: "POST", path: "/users", controller: "UserController", action: "store", file: "routes/api.php", line: 1}
	ep := buildOneEndpoint(dir, store, parsed)
	if ep.Request == nil || ep.Request.Body == nil {
		t.Fatalf("expected request body, got %+v", ep.Request)
	}
	if len(ep.Responses) == 0 {
		t.Fatalf("expected responses, got %+v", ep.Responses)
	}

	show := routeInfo{method: "GET", path: "/users/{id}", controller: "UserController", action: "show", file: "routes/api.php", line: 2}
	ep2 := buildOneEndpoint(dir, show, parsed)
	if ep2.Request == nil || len(ep2.Request.PathParams) == 0 {
		t.Fatalf("expected path params, got %+v", ep2.Request)
	}
}

func TestBuildEndpoints_Multiple(t *testing.T) {
	routes := []routeInfo{
		{method: "GET", path: "/a"},
		{method: "POST", path: "/b"},
	}
	eps := buildEndpoints(t.TempDir(), routes, map[string]*fileInfo{})
	if len(eps) != 2 {
		t.Fatalf("expected 2, got %d", len(eps))
	}
}

func TestBuildRequest_Nil(t *testing.T) {
	// no path params, no form request -> nil
	if r := buildRequest(t.TempDir(), nil, nil, map[string]*fileInfo{}); r != nil {
		t.Fatalf("expected nil, got %+v", r)
	}
}

func TestBuildRequest_PathParamsOnly(t *testing.T) {
	pp := []scanner.Param{{Name: "id", Type: "integer"}}
	r := buildRequest(t.TempDir(), pp, nil, map[string]*fileInfo{})
	if r == nil || len(r.PathParams) != 1 || r.Body != nil {
		t.Fatalf("got %+v", r)
	}
}

func TestBuildRequest_WithFormRequest(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "app/Http/Requests/StoreReq.php", `<?php
namespace App\Http\Requests;
use Illuminate\Foundation\Http\FormRequest;
class StoreReq extends FormRequest {
    public function rules(): array { return [ 'name' => 'required|string' ]; }
}
`)
	cm := &controllerMethod{formRequestRef: "StoreReq"}
	r := buildRequest(dir, nil, cm, map[string]*fileInfo{})
	if r == nil || r.Body == nil || len(r.Body.Fields) != 1 {
		t.Fatalf("got %+v", r)
	}
}

//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestBuildOneEndpoint_WithController 테스트
package laravel

import "testing"

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

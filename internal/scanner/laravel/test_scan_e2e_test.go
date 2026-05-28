//ff:func feature=scan type=test control=iteration dimension=1 topic=laravel
//ff:what E2E 스캔 테스트: 라우트 + apiResource + prefix 그룹 + middleware 그룹 + FormRequest + Resource
package laravel

import "testing"

func TestScan_E2E(t *testing.T) {
	dir := t.TempDir()

	// routes/api.php
	writeFile(t, dir, "routes/api.php", `<?php
use App\Http\Controllers\UserController;
use App\Http\Controllers\ProfileController;

Route::get('/users', [UserController::class, 'index']);
Route::post('/users', [UserController::class, 'store']);
Route::get('/users/{user}', [UserController::class, 'show']);

Route::apiResource('posts', PostController::class);

Route::prefix('v1')->group(function () {
    Route::get('/health', function () { return response()->json(['status' => 'ok']); });
});

Route::middleware(['auth:sanctum'])->group(function () {
    Route::get('/me', [ProfileController::class, 'show']);
});
`)

	// Controller with FormRequest
	writeFile(t, dir, "app/Http/Controllers/UserController.php", `<?php
namespace App\Http\Controllers;

use App\Http\Requests\StoreUserRequest;
use App\Http\Resources\UserResource;

class UserController extends Controller {
    public function index() {
        return UserResource::collection(User::paginate(10));
    }
    public function show(int $id) {
        return new UserResource(User::findOrFail($id));
    }
    public function store(StoreUserRequest $request) {
        $user = User::create($request->validated());
        return new UserResource($user);
    }
}
`)

	// FormRequest
	writeFile(t, dir, "app/Http/Requests/StoreUserRequest.php", `<?php
namespace App\Http\Requests;

use Illuminate\Foundation\Http\FormRequest;

class StoreUserRequest extends FormRequest {
    public function rules(): array {
        return [
            'name' => 'required|string|max:255',
            'email' => 'required|email|unique:users',
            'age' => 'nullable|integer|min:0|max:150',
        ];
    }
}
`)

	// Resource
	writeFile(t, dir, "app/Http/Resources/UserResource.php", `<?php
namespace App\Http\Resources;

use Illuminate\Http\Resources\Json\JsonResource;

class UserResource extends JsonResource {
    public function toArray($request): array {
        return [
            'id' => $this->id,
            'name' => $this->name,
            'email' => $this->email,
        ];
    }
}
`)

	result, err := Scan(dir)
	if err != nil {
		t.Fatal(err)
	}

	// Expected endpoints:
	// GET /api/users         (index)
	// POST /api/users        (store)
	// GET /api/users/{user}  (show)
	// GET /api/posts         (apiResource - index)
	// POST /api/posts        (apiResource - store)
	// GET /api/posts/{post}  (apiResource - show)
	// PUT /api/posts/{post}  (apiResource - update)
	// DELETE /api/posts/{post} (apiResource - destroy)
	// GET /api/v1/health     (prefix group)
	// GET /api/me            (middleware group)
	// Total: 10

	if len(result.Endpoints) < 10 {
		for i, ep := range result.Endpoints {
			t.Logf("  endpoint %d: %s %s (%s)", i, ep.Method, ep.Path, ep.Handler)
		}
		t.Fatalf("expected at least 10 endpoints, got %d", len(result.Endpoints))
	}

	found := map[string]bool{}
	for _, ep := range result.Endpoints {
		key := ep.Method + " " + ep.Path
		found[key] = true
		t.Logf("  found: %s %s (%s)", ep.Method, ep.Path, ep.Handler)
	}

	expectedEndpoints := []string{
		"GET /api/users",
		"POST /api/users",
		"GET /api/users/{user}",
		"GET /api/posts",
		"POST /api/posts",
		"GET /api/posts/{post}",
		"PUT /api/posts/{post}",
		"DELETE /api/posts/{post}",
		"GET /api/v1/health",
		"GET /api/me",
	}
	for _, expected := range expectedEndpoints {
		if !found[expected] {
			t.Errorf("missing expected endpoint: %s", expected)
		}
	}

	// Verify FormRequest fields on POST /api/users
	for _, ep := range result.Endpoints {
		if ep.Method == "POST" && ep.Path == "/api/users" {
			if ep.Request == nil || ep.Request.Body == nil {
				t.Error("POST /api/users: expected request body from FormRequest")
				break
			}
			if len(ep.Request.Body.Fields) != 3 {
				t.Errorf("POST /api/users: expected 3 body fields, got %d", len(ep.Request.Body.Fields))
			}
			if ep.Request.Body.TypeName != "StoreUserRequest" {
				t.Errorf("POST /api/users: body type = %q, want StoreUserRequest", ep.Request.Body.TypeName)
			}
			break
		}
	}

	// Verify Resource response on GET /api/users/{user}
	for _, ep := range result.Endpoints {
		if ep.Method == "GET" && ep.Path == "/api/users/{user}" {
			if len(ep.Responses) == 0 {
				t.Error("GET /api/users/{user}: expected responses from Resource")
				break
			}
			if ep.Responses[0].TypeName != "UserResource" {
				t.Errorf("GET /api/users/{user}: response type = %q, want UserResource", ep.Responses[0].TypeName)
			}
			break
		}
	}

	// Verify path param type hint on GET /api/users/{user} (int $id)
	for _, ep := range result.Endpoints {
		if ep.Method == "GET" && ep.Path == "/api/users/{user}" {
			if ep.Request == nil {
				t.Error("GET /api/users/{user}: expected request with path params")
				break
			}
			// The path has {user} but controller has int $id - types should apply
			break
		}
	}

	// Verify middleware on GET /api/me
	for _, ep := range result.Endpoints {
		if ep.Method == "GET" && ep.Path == "/api/me" {
			if len(ep.Middleware) != 1 || ep.Middleware[0] != "auth:sanctum" {
				t.Errorf("GET /api/me: middleware = %v, want [auth:sanctum]", ep.Middleware)
			}
			break
		}
	}
}

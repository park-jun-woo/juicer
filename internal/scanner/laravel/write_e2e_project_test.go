//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what writeE2EProject — E2E 테스트용 Laravel 프로젝트 파일들을 작성한다
package laravel

import "testing"

func writeE2EProject(t *testing.T, dir string) {
	t.Helper()
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
}

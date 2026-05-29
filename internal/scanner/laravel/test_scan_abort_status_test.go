//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestScan_AbortAndConstantStatus -- abort() 응답 추가 + Response::HTTP_* 상수 상태 코드 매핑
package laravel

import "testing"

func TestScan_AbortAndConstantStatus(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "routes/api.php", `<?php
use App\Http\Controllers\OrderController;

Route::get('/orders/{order}', [OrderController::class, 'show']);
Route::post('/orders', [OrderController::class, 'store']);
`)
	writeFile(t, dir, "app/Http/Controllers/OrderController.php", `<?php
namespace App\Http\Controllers;

use Symfony\Component\HttpFoundation\Response;

class OrderController extends Controller {
    public function show(int $id) {
        $order = Order::find($id);
        abort_if($order === null, 404);
        return response()->json($order);
    }
    public function store() {
        return response()->json($data, Response::HTTP_CREATED);
    }
}
`)

	result, err := Scan(dir)
	if err != nil {
		t.Fatal(err)
	}

	show := findEndpoint(result.Endpoints, "GET", "/api/orders/{order}")
	if show == nil {
		t.Fatalf("show endpoint not found: %+v", result.Endpoints)
	}
	if !hasStatus(show.Responses, "404") {
		t.Errorf("show: expected 404 from abort_if, got %+v", show.Responses)
	}
	if !hasStatus(show.Responses, "200") {
		t.Errorf("show: expected 200 from json(), got %+v", show.Responses)
	}

	store := findEndpoint(result.Endpoints, "POST", "/api/orders")
	if store == nil {
		t.Fatalf("store endpoint not found: %+v", result.Endpoints)
	}
	if !hasStatus(store.Responses, "201") {
		t.Errorf("store: expected 201 from Response::HTTP_CREATED, got %+v", store.Responses)
	}
}

//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestBuildRequest_WithFormRequest 테스트
package laravel

import "testing"

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

//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestBuildFormRequestBody_WithFields 테스트
package laravel

import "testing"

func TestBuildFormRequestBody_WithFields(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "app/Http/Requests/StoreUserRequest.php", `<?php
namespace App\Http\Requests;
use Illuminate\Foundation\Http\FormRequest;
class StoreUserRequest extends FormRequest {
    public function rules(): array {
        return [ 'name' => 'required|string' ];
    }
}
`)
	cm := &controllerMethod{formRequestRef: "StoreUserRequest"}
	b := buildFormRequestBody(dir, cm, map[string]*fileInfo{})
	if b == nil {
		t.Fatal("expected non-nil body")
	}
	if b.VarName != "request" || b.Method != "json" || b.TypeName != "StoreUserRequest" {
		t.Fatalf("meta: %+v", b)
	}
	if len(b.Fields) != 1 || b.Fields[0].Name != "name" {
		t.Fatalf("fields: %+v", b.Fields)
	}
}

//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what buildParamTypeMap / buildFormRequestBody 테스트
package laravel

import "testing"

func TestBuildParamTypeMap(t *testing.T) {
	cm := &controllerMethod{params: []methodParam{
		{name: "id", typeName: "int"},
		{name: "name", typeName: "string"},
		{name: "request", typeName: "Request"}, // skipped by name
		{name: "untyped", typeName: ""},        // skipped by empty type
		{name: "weird", typeName: "SomeClass"}, // skipped: phpTypeToOpenAPI empty
	}}
	m := buildParamTypeMap(cm)
	if m["id"] != "integer" || m["name"] != "string" {
		t.Fatalf("got %v", m)
	}
	if _, ok := m["request"]; ok {
		t.Fatal("request should be skipped")
	}
	if _, ok := m["untyped"]; ok {
		t.Fatal("untyped should be skipped")
	}
	if _, ok := m["weird"]; ok {
		t.Fatal("unmapped type should be skipped")
	}
}

func TestBuildFormRequestBody_NoFields(t *testing.T) {
	// formRequestRef that resolves to nothing -> nil body
	cm := &controllerMethod{formRequestRef: "NonexistentRequest"}
	if b := buildFormRequestBody(t.TempDir(), cm, map[string]*fileInfo{}); b != nil {
		t.Fatalf("expected nil, got %+v", b)
	}
}

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

//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestApplyParamDecorators_All 테스트
package nestjs

import "testing"

func TestApplyParamDecorators_All(t *testing.T) {
	result := &methodParams{}
	decorators := []decoratorInfo{
		{name: DecParam, arg: "id"},
		{name: DecQuery, arg: "page"},
		{name: DecBody, arg: ""},
		{name: DecUploadedFile, arg: "file"},
	}
	applyParamDecorators(decorators, "default", "string", result)
	if len(result.pathParams) != 1 || result.pathParams[0].Name != "id" {
		t.Fatal("expected path param id")
	}
	if len(result.queryParams) != 1 || result.queryParams[0].Name != "page" {
		t.Fatal("expected query param page")
	}
	if result.bodyType != "string" {
		t.Fatal("expected body type string")
	}
	if len(result.files) != 1 || result.files[0].Name != "file" {
		t.Fatal("expected file param")
	}
}

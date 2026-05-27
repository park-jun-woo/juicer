//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestSubstituteTypes 제네릭 타입 치환 테스트
package nestjs

import "testing"

func TestSubstituteTypes(t *testing.T) {
	eps := []endpointInfo{
		{returnType: "D", bodyType: "CreateDtoType", queryDTOType: ""},
		{returnType: "ResponsePaginateDto<B>", bodyType: "", queryDTOType: "QueryDto<B>"},
		{returnType: "", bodyType: "", queryDTOType: ""},
	}
	typeMap := map[string]string{
		"D":             "Category",
		"B":             "CategoryDto",
		"CreateDtoType": "CreateCategoryDto",
	}
	substituteTypes(eps, typeMap)

	if eps[0].returnType != "Category" {
		t.Errorf("eps[0].returnType = %q, want Category", eps[0].returnType)
	}
	if eps[0].bodyType != "CreateCategoryDto" {
		t.Errorf("eps[0].bodyType = %q, want CreateCategoryDto", eps[0].bodyType)
	}
	if eps[1].returnType != "ResponsePaginateDto<CategoryDto>" {
		t.Errorf("eps[1].returnType = %q, want ResponsePaginateDto<CategoryDto>", eps[1].returnType)
	}
	if eps[1].queryDTOType != "QueryDto<CategoryDto>" {
		t.Errorf("eps[1].queryDTOType = %q, want QueryDto<CategoryDto>", eps[1].queryDTOType)
	}
	if eps[2].returnType != "" {
		t.Errorf("eps[2].returnType = %q, want empty", eps[2].returnType)
	}
}

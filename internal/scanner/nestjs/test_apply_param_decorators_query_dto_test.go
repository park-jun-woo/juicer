//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestApplyParamDecorators_QueryDTO 테스트
package nestjs

import "testing"

func TestApplyParamDecorators_QueryDTO(t *testing.T) {
	result := &methodParams{}
	decorators := []decoratorInfo{
		{name: DecQuery, arg: ""},
	}
	applyParamDecorators(decorators, "dto", "ListUserReqDto", result)
	if result.queryDTOType != "ListUserReqDto" {
		t.Fatalf("expected queryDTOType=ListUserReqDto, got %q", result.queryDTOType)
	}
	if len(result.queryParams) != 0 {
		t.Fatalf("expected 0 queryParams, got %d", len(result.queryParams))
	}
}

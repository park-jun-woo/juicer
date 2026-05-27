//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestApplyParamDecorators_DefaultName 테스트
package nestjs

import "testing"

func TestApplyParamDecorators_DefaultName(t *testing.T) {
	result := &methodParams{}
	decorators := []decoratorInfo{
		{name: DecParam, arg: ""},
	}
	applyParamDecorators(decorators, "myParam", "number", "", result)
	if result.pathParams[0].Name != "myParam" {
		t.Fatalf("expected myParam, got %s", result.pathParams[0].Name)
	}
}

//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what pyTypeToOpenAPI: Annotated 위임 / 커스텀 모델 default(object)
package fastapi

import "testing"

func TestPyTypeToOpenAPI_Annotated(t *testing.T) {
	// Annotated[int, ...] -> tryAnnotated unwraps to inner type
	oa := pyTypeToOpenAPI("Annotated[int, Field()]")
	if oa.Type != "integer" {
		t.Fatalf("got %+v", oa)
	}
}

func TestPyTypeToOpenAPI_CustomModel(t *testing.T) {
	// unknown custom type -> default case -> object
	oa := pyTypeToOpenAPI("UserResponse")
	if oa.Type != "object" {
		t.Fatalf("got %+v", oa)
	}
}

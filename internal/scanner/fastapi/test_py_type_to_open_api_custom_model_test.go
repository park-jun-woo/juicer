//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestPyTypeToOpenAPI_CustomModel 테스트
package fastapi

import "testing"

func TestPyTypeToOpenAPI_CustomModel(t *testing.T) {

	oa := pyTypeToOpenAPI("UserResponse")
	if oa.Type != "object" {
		t.Fatalf("got %+v", oa)
	}
}

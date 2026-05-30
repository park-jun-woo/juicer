//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestPyTypeToOpenAPI_Annotated 테스트
package fastapi

import "testing"

func TestPyTypeToOpenAPI_Annotated(t *testing.T) {

	oa := pyTypeToOpenAPI("Annotated[int, Field()]")
	if oa.Type != "integer" {
		t.Fatalf("got %+v", oa)
	}
}

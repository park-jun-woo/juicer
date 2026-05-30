//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what TestClassifyFormParam 테스트
package dotnet

import "testing"

func TestClassifyFormParam(t *testing.T) {
	ep := &endpointInfo{}
	classifyFormParam(ep, "IFormFile", "file")
	if len(ep.files) != 1 {
		t.Fatalf("file: %+v", ep.files)
	}
	ep2 := &endpointInfo{}
	classifyFormParam(ep2, "string", "name")
	if len(ep2.formFields) != 1 {
		t.Fatalf("form: %+v", ep2.formFields)
	}
}

//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what TestCsharpTypeToOpenAPIType_Round5 테스트
package dotnet

import "testing"

func TestCsharpTypeToOpenAPIType_Round5(t *testing.T) {
	if got := csharpTypeToOpenAPIType("int"); got != "integer" {
		t.Errorf("int: got %q", got)
	}

	if got := csharpTypeToOpenAPIType("void"); got != "string" {
		t.Errorf("void fallback: got %q", got)
	}
}

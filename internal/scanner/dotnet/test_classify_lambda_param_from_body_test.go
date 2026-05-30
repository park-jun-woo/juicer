//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what TestClassifyLambdaParam_FromBody 테스트
package dotnet

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestClassifyLambdaParam_FromBody(t *testing.T) {
	root, src := parseCS(t, `class C { void M() { app.MapPost("/x", ([FromBody] UserDto dto) => Results.Ok()); } }`)
	params := findAllByType(root, "parameter")
	var req scanner.Request
	classifyLambdaParam(params[0], src, "/x", &req)
	if req.Body == nil || req.Body.TypeName != "UserDto" {
		t.Fatalf("got %+v", req.Body)
	}
}

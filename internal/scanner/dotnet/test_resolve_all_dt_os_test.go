//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what TestResolveAllDTOs 테스트
package dotnet

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestResolveAllDTOs(t *testing.T) {
	fi := csFileInfo(t, `class UserDto { public string Name { get; set; } }`)
	endpoints := []scanner.Endpoint{{Responses: []scanner.Response{{Status: "200"}}}}
	dr := dtoRequest{typeName: "UserDto", epIdx: 0}
	resolveAllDTOs([]dtoRequest{dr}, endpoints, []*fileInfo{fi})
	if len(endpoints[0].Responses[0].Fields) != 1 {
		t.Fatalf("got %+v", endpoints[0].Responses)
	}
}

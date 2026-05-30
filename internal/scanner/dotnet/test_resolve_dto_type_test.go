//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what TestResolveDTOType 테스트
package dotnet

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestResolveDTOType(t *testing.T) {
	fi := csFileInfo(t, `class UserDto { public string Name { get; set; } public int Age { get; set; } }`)
	dr := dtoRequest{typeName: "UserDto"}
	fields := resolveDTOType(dr, []*fileInfo{fi}, map[string][]scanner.Field{})
	if len(fields) != 2 || fields[0].Name != "Name" {
		t.Fatalf("got %+v", fields)
	}
}

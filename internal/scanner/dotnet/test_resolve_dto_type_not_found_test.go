//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what TestResolveDTOType_NotFound 테스트
package dotnet

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestResolveDTOType_NotFound(t *testing.T) {
	fi := csFileInfo(t, `class Other {}`)
	if fields := resolveDTOType(dtoRequest{typeName: "Missing"}, []*fileInfo{fi}, map[string][]scanner.Field{}); fields != nil {
		t.Fatalf("got %+v", fields)
	}
}

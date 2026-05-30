//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what TestResolveDTOType_Cached 테스트
package dotnet

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestResolveDTOType_Cached(t *testing.T) {
	cache := map[string][]scanner.Field{"X": {{Name: "a"}}}
	fields := resolveDTOType(dtoRequest{typeName: "X"}, nil, cache)
	if len(fields) != 1 {
		t.Fatalf("got %+v", fields)
	}
}

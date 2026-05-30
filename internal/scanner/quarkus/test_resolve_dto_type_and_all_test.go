//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestResolveDTOTypeAndAll 테스트
package quarkus

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"path/filepath"
	"testing"
)

func TestResolveDTOTypeAndAll(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "UserDto.java", `class UserDto { private String name; }`)
	dr := dtoRequest{
		typeName:    "UserDto",
		imports:     map[string]string{},
		referrer:    filepath.Join(dir, "Resource.java"),
		projectRoot: dir,
		epIdx:       0,
	}
	writeFile(t, dir, "Resource.java", `@Path("/x") class Resource {}`)
	cache := map[string][]scanner.Field{}
	fields := resolveDTOType(dr, dir, cache)
	if len(fields) != 1 || fields[0].Name != "name" {
		t.Fatalf("got %+v", fields)
	}

	endpoints := []scanner.Endpoint{{Responses: []scanner.Response{{Status: "200"}}}}
	resolveAllDTOs([]dtoRequest{dr}, endpoints, dir)
	if len(endpoints[0].Responses[0].Fields) != 1 {
		t.Fatalf("resolveAllDTOs did not assign: %+v", endpoints[0].Responses)
	}
}

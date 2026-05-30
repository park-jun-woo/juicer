//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestResolveDTOTypeAndAll 테스트
package spring

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"path/filepath"
	"testing"
)

func TestResolveDTOTypeAndAll(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "UserDto.java", `class UserDto { private String name; }`)
	writeFile(t, dir, "C.java", `@RestController class C {}`)
	dr := dtoRequest{
		typeName:    "UserDto",
		imports:     map[string]string{},
		referrer:    filepath.Join(dir, "C.java"),
		projectRoot: dir,
		epIdx:       0,
	}
	fields := resolveDTOType(dr, dir, map[string][]scanner.Field{})
	if len(fields) != 1 || fields[0].Name != "name" {
		t.Fatalf("got %+v", fields)
	}
	endpoints := []scanner.Endpoint{{Responses: []scanner.Response{{Status: "200"}}}}
	resolveAllDTOs([]dtoRequest{dr}, endpoints, dir)
	if len(endpoints[0].Responses[0].Fields) != 1 {
		t.Fatalf("resolveAllDTOs: %+v", endpoints[0].Responses)
	}
}

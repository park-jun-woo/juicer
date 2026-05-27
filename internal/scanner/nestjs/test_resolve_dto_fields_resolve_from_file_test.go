//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestResolveDTOFields_ResolveFromFile 테스트
package nestjs

import (
	"testing"
	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestResolveDTOFields_ResolveFromFile(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "dto/user.dto.ts", `
export class UserDto {
  name: string;
  email: string;
}
`)
	cache := make(map[string][]scanner.Field)
	dr := dtoRequest{
		typeName: "UserDto",
		imports:  map[string]string{"UserDto": "./dto/user.dto"},
		referrer: dir + "/src/controller.ts",
	}
	// referrer is in dir/src/, import resolves to dir/dto/user.dto.ts
	// But referrerDir is dir/src, and resolveImportPath joins dir/src + ./dto/user.dto
	// That would be dir/src/dto/user.dto.ts which doesn't exist.
	// Need to adjust: put the DTO relative to referrer dir.
	writeFile(t, dir, "src/dto/user.dto.ts", `
export class UserDto {
  name: string;
  email: string;
}
`)
	fields, err := resolveDTOFields(dr, cache)
	if err != nil {
		t.Fatal(err)
	}
	if len(fields) != 2 {
		t.Fatalf("expected 2 fields, got %d", len(fields))
	}
}

//ff:func feature=scan type=test control=iteration dimension=1 topic=nestjs
//ff:what TestRegisterNestedSchemas — 중첩 DTO/enum이 별도 스키마로 재귀 등록된다
package nestjs

import (
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestRegisterNestedSchemas(t *testing.T) {
	dir := t.TempDir()

	writeFile(t, dir, "dto/role.enum.ts", `
export enum AlbumUserRole { Editor = 'editor', Viewer = 'viewer' }
`)
	writeFile(t, dir, "dto/album-user.dto.ts", `
import { AlbumUserRole } from './role.enum';
export class AlbumUserResponseDto {
  userId: string;
  role: AlbumUserRole;
}
`)
	writeFile(t, dir, "dto/album.dto.ts", `
import { AlbumUserResponseDto } from './album-user.dto';
export class AlbumResponseDto {
  id: string;
  owner: AlbumUserResponseDto;
  albumUsers: AlbumUserResponseDto[];
}
`)

	imports := map[string]string{"AlbumResponseDto": "./dto/album.dto"}
	dr := dtoRequest{
		typeName: "AlbumResponseDto", imports: imports,
		referrer: dir + "/x.controller.ts", projectRoot: dir,
		epIdx: 0, isBody: false,
	}
	cache := make(map[string][]scanner.Field)
	// Resolve top-level so cache holds its fields (with Refs).
	top, err := resolveDTOFields(dr, cache)
	if err != nil {
		t.Fatal(err)
	}
	if len(top) == 0 {
		t.Fatal("expected top-level fields")
	}

	schemas := make(map[string]any)
	registerNestedSchemas([]dtoRequest{dr}, cache, schemas)

	// Nested DTO and enum must be registered as separate schemas.
	if _, ok := schemas["AlbumUserResponseDto"]; !ok {
		t.Fatalf("nested DTO not registered: %v keys", keysOf(schemas))
	}
	enum, ok := schemas["AlbumUserRole"].(map[string]any)
	if !ok {
		t.Fatalf("nested enum not registered: %v", keysOf(schemas))
	}
	if enum["type"] != "string" || enum["enum"] == nil {
		t.Fatalf("enum schema malformed: %v", enum)
	}

	// The top-level AlbumResponseDto field "owner" must carry a $ref.
	var ownerRef string
	for _, f := range top {
		if f.Name == "owner" {
			ownerRef = f.Ref
		}
	}
	if ownerRef != "AlbumUserResponseDto" {
		t.Fatalf("owner field should ref AlbumUserResponseDto, got %q", ownerRef)
	}
}

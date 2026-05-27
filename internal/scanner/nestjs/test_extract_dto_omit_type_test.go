//ff:func feature=scan type=test control=iteration dimension=1 topic=nestjs
//ff:what TestExtractDTO_OmitType 테스트
package nestjs

import (
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestExtractDTO_OmitType(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "dto/create-user.dto.ts", `
export class CreateUserReqDto {
  name: string;
  email: string;
  password: string;
}
`)
	writeFile(t, dir, "dto/update-user.dto.ts", `
import { OmitType } from '@nestjs/mapped-types';
import { CreateUserReqDto } from './create-user.dto';
export class UpdateUserReqDto extends OmitType(CreateUserReqDto, ['password']) {}
`)
	cache := make(map[string][]scanner.Field)
	imports := map[string]string{"CreateUserReqDto": "./create-user.dto"}
	fields, err := extractDTO(dir+"/dto/update-user.dto.ts", "UpdateUserReqDto", imports, "", cache)
	if err != nil {
		t.Fatal(err)
	}
	if len(fields) != 2 {
		t.Fatalf("expected 2 fields (name, email), got %d", len(fields))
	}
	for _, f := range fields {
		if f.Name == "password" {
			t.Fatal("password should be omitted")
		}
	}
}

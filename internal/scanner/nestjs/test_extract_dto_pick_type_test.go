//ff:func feature=scan type=test control=iteration dimension=1 topic=nestjs
//ff:what TestExtractDTO_PickType 테스트
package nestjs

import (
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestExtractDTO_PickType(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "dto/base.dto.ts", `
export class BaseDto {
  name: string;
  email: string;
  age: number;
}
`)
	writeFile(t, dir, "dto/name-email.dto.ts", `
import { PickType } from '@nestjs/mapped-types';
import { BaseDto } from './base.dto';
export class NameEmailDto extends PickType(BaseDto, ['name', 'email']) {}
`)
	cache := make(map[string][]scanner.Field)
	imports := map[string]string{"BaseDto": "./base.dto"}
	fields, err := extractDTO(dir+"/dto/name-email.dto.ts", "NameEmailDto", imports, "", cache)
	if err != nil {
		t.Fatal(err)
	}
	if len(fields) != 2 {
		t.Fatalf("expected 2 fields (name, email), got %d", len(fields))
	}
	names := make(map[string]bool)
	for _, f := range fields {
		names[f.Name] = true
	}
	if !names["name"] || !names["email"] {
		t.Fatal("expected name and email fields")
	}
}

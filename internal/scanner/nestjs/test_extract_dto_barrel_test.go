//ff:func feature=scan type=test control=iteration dimension=1 topic=nestjs
//ff:what TestExtractDTO_Barrel barrel index.ts 경유 DTO 추출 테스트
package nestjs

import (
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestExtractDTO_Barrel(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "dtos/create-category.dto.ts", `
export class CreateCategoryDto {
  name: string;
  quantity: number;
  description: string;
}
`)
	writeFile(t, dir, "dtos/index.ts", `
export * from './create-category.dto';
`)

	cache := make(map[string][]scanner.Field)
	fields, err := extractDTO(dir+"/dtos/index.ts", "CreateCategoryDto", nil, "", cache)
	if err != nil {
		t.Fatal(err)
	}
	if len(fields) != 3 {
		t.Fatalf("expected 3 fields, got %d", len(fields))
	}
	names := make(map[string]bool)
	for _, f := range fields {
		names[f.Name] = true
	}
	for _, want := range []string{"name", "quantity", "description"} {
		if !names[want] {
			t.Fatalf("missing field %q", want)
		}
	}
}

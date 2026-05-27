//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestExtractDTO_PartialType 테스트
package nestjs

import (
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestExtractDTO_PartialType(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "dto/create-task.dto.ts", `
export class CreateTaskDto {
  title: string;
  description: string;
}
`)
	writeFile(t, dir, "dto/update-task.dto.ts", `
import { PartialType } from '@nestjs/mapped-types';
import { CreateTaskDto } from './create-task.dto';
export class UpdateTaskDto extends PartialType(CreateTaskDto) {}
`)
	cache := make(map[string][]scanner.Field)
	imports := map[string]string{"CreateTaskDto": "./create-task.dto"}
	fields, err := extractDTO(dir+"/dto/update-task.dto.ts", "UpdateTaskDto", imports, "", cache)
	if err != nil {
		t.Fatal(err)
	}
	if len(fields) != 2 {
		t.Fatalf("expected 2 fields, got %d", len(fields))
	}
}

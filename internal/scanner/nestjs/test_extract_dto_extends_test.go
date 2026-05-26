//ff:func feature=scan type=test control=iteration dimension=1 topic=nestjs
//ff:what TestExtractDTO_Extends 테스트
package nestjs

import (
	"testing"

	"github.com/park-jun-woo/juicer/internal/scanner"
)

func TestExtractDTO_Extends(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "dto/create-post.dto.ts", `
export class CreatePostReqDto {
  title: string;
  content: string;
}
`)
	writeFile(t, dir, "dto/update-post.dto.ts", `
import { CreatePostReqDto } from './create-post.dto';
export class UpdatePostReqDto extends CreatePostReqDto {
  id: number;
}
`)
	cache := make(map[string][]scanner.Field)
	imports := map[string]string{"CreatePostReqDto": "./create-post.dto"}
	fields, err := extractDTO(dir+"/dto/update-post.dto.ts", "UpdatePostReqDto", imports, "", cache)
	if err != nil {
		t.Fatal(err)
	}
	if len(fields) != 3 {
		t.Fatalf("expected 3 fields (title, content, id), got %d", len(fields))
	}
	names := make(map[string]bool)
	for _, f := range fields {
		names[f.Name] = true
	}
	for _, want := range []string{"title", "content", "id"} {
		if !names[want] {
			t.Fatalf("missing field %q", want)
		}
	}
}

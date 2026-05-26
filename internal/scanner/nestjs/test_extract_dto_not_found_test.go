//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestExtractDTO_NotFound 테스트
package nestjs

import "testing"

func TestExtractDTO_NotFound(t *testing.T) {
	dir := t.TempDir()
	dto := `export class OtherDto { name: string; }`
	writeFile(t, dir, "other.dto.ts", dto)
	fields, err := extractDTO(dir+"/other.dto.ts", "CreateUserDto")
	if err != nil {
		t.Fatal(err)
	}
	if fields != nil {
		t.Fatal("expected nil for not found class")
	}
}

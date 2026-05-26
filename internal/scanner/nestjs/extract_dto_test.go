//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestExtractDTO_Found 테스트
package nestjs

import "testing"

func TestExtractDTO_Found(t *testing.T) {
	dir := t.TempDir()
	dto := `
export class CreateUserDto {
  name: string;
  email: string;
}
`
	writeFile(t, dir, "create-user.dto.ts", dto)
	fields, err := extractDTO(dir+"/create-user.dto.ts", "CreateUserDto")
	if err != nil {
		t.Fatal(err)
	}
	if len(fields) != 2 {
		t.Fatalf("expected 2 fields, got %d", len(fields))
	}
}

//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what extractClassProperties 테스트
package nestjs

import "testing"

func TestExtractClassProperties_Basic(t *testing.T) {
	src := []byte(`
export class CreateUserDto {
  name: string;
  email: string;
}
`)
	root, err := parseTypeScript(src)
	if err != nil {
		t.Fatal(err)
	}
	classes := findAllByType(root, "class_declaration")
	if len(classes) == 0 {
		t.Fatal("no classes")
	}
	fields := extractClassProperties(classes[0], src)
	if len(fields) != 2 {
		t.Fatalf("expected 2 fields, got %d", len(fields))
	}
}

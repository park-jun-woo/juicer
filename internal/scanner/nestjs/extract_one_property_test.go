//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what extractOneProperty 테스트
package nestjs

import "testing"

func TestExtractOneProperty_Basic(t *testing.T) {
	src := []byte(`
export class CreateUserDto {
  name: string;
  age?: number;
}
`)
	root, err := parseTypeScript(src)
	if err != nil {
		t.Fatal(err)
	}
	props := findAllByType(root, "public_field_definition")
	if len(props) < 2 {
		t.Fatalf("expected 2 properties, got %d", len(props))
	}
	f0 := extractOneProperty(props[0], src)
	if f0.name != "name" || f0.tsType != "string" {
		t.Fatalf("unexpected field 0: %+v", f0)
	}
	f1 := extractOneProperty(props[1], src)
	if f1.name != "age" || !f1.optional {
		t.Fatalf("unexpected field 1: %+v", f1)
	}
}

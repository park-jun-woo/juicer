//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestExtractEnumFromApiProperty_NoApiProperty 테스트
package nestjs

import "testing"

func TestExtractEnumFromApiProperty_NoApiProperty(t *testing.T) {
	src := []byte(`
class Dto {
  @IsString()
  name: string;
}
`)
	root, _ := parseTypeScript(src)
	props := findAllByType(root, "public_field_definition")
	if len(props) == 0 {
		t.Skip("no field def")
	}
	if vals := extractEnumFromApiProperty(props[0], src); vals != nil {
		t.Fatalf("expected nil, got %v", vals)
	}
}

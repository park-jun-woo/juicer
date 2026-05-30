//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestExtractEnumFromApiProperty 테스트
package nestjs

import "testing"

func TestExtractEnumFromApiProperty(t *testing.T) {
	src := []byte(`
class Dto {
  @ApiProperty({ enum: ['a', 'b'] })
  status: string;
}
`)
	root, _ := parseTypeScript(src)
	props := findAllByType(root, "public_field_definition")
	if len(props) == 0 {
		t.Skip("no field def")
	}
	vals := extractEnumFromApiProperty(props[0], src)
	if len(vals) != 2 || vals[0] != "a" {
		t.Fatalf("got %v", vals)
	}
}

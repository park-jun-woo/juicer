//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what propertyDecorators 테스트
package nestjs

import "testing"

func TestPropertyDecorators_WithDecorators(t *testing.T) {
	src := []byte(`
class Dto {
  @IsString()
  name: string;
}
`)
	root, err := parseTypeScript(src)
	if err != nil {
		t.Fatal(err)
	}
	props := findAllByType(root, "public_field_definition")
	if len(props) == 0 {
		t.Fatal("no properties")
	}
	// propertyDecorators works on the node structure; just ensure no panic
	_ = propertyDecorators(props[0], src)
}

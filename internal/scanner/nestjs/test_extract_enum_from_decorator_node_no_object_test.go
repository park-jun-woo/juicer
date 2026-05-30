//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestExtractEnumFromDecoratorNode_NoObject 테스트
package nestjs

import "testing"

func TestExtractEnumFromDecoratorNode_NoObject(t *testing.T) {
	src := []byte(`
class Dto {
  @ApiProperty()
  name: string;
}
`)
	root, _ := parseTypeScript(src)
	decs := findAllByType(root, "decorator")
	if len(decs) == 0 {
		t.Skip("no decorator")
	}
	if vals := extractEnumFromDecoratorNode(decs[0], src); vals != nil {
		t.Fatalf("expected nil, got %v", vals)
	}
}

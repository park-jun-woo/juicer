//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what extractEnumFromApiProperty / extractEnumFromDecoratorNode / extractFactoryBaseClass 테스트
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

func TestExtractFactoryBaseClass(t *testing.T) {
	src := []byte(`const x = PartialType(CreateTaskDto);`)
	root, _ := parseTypeScript(src)
	args := findAllByType(root, "arguments")[0]
	if got := extractFactoryBaseClass(args, src); got != "CreateTaskDto" {
		t.Fatalf("got %q", got)
	}
}

func TestExtractFactoryBaseClass_None(t *testing.T) {
	src := []byte(`const x = foo('str');`)
	root, _ := parseTypeScript(src)
	args := findAllByType(root, "arguments")[0]
	if got := extractFactoryBaseClass(args, src); got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}

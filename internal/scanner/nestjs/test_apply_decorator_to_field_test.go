//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestApplyDecoratorToField 테스트
package nestjs

import "testing"

func TestApplyDecoratorToField(t *testing.T) {
	src := []byte(`class D { @IsOptional() name: string; }`)
	root, _ := parseTypeScript(src)
	props := findAllByType(root, "public_field_definition")
	var f dtoField
	applyDecoratorToField(decoratorInfo{name: "IsOptional"}, nil, src, &f)
	if !f.optional {
		t.Fatal("expected optional")
	}
	var f2 dtoField
	applyDecoratorToField(decoratorInfo{name: "MinLength", arg: "5"}, nil, src, &f2)
	if f2.minLength == nil || *f2.minLength != 5 {
		t.Fatalf("minLength: %v", f2.minLength)
	}
	var f3 dtoField
	applyDecoratorToField(decoratorInfo{name: "IsEnum", arg: "Status"}, nil, src, &f3)
	if f3.enumTypeName != "Status" {
		t.Fatalf("enumTypeName: %q", f3.enumTypeName)
	}
	_ = props
}

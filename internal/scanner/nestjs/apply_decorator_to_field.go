//ff:func feature=scan type=extract control=selection topic=nestjs
//ff:what 단일 데코레이터 정보를 dtoField에 적용한다
package nestjs

import sitter "github.com/smacker/go-tree-sitter"

// applyDecoratorToField applies a single decorator's info to the dtoField.
func applyDecoratorToField(d decoratorInfo, prop *sitter.Node, src []byte, f *dtoField) {
	f.validators = append(f.validators, d.name)
	switch d.name {
	case "IsOptional":
		f.optional = true
	case "ApiProperty":
		f.enum = extractEnumFromApiProperty(prop, src)
	case "IsEnum":
		if d.arg != "" {
			f.enumTypeName = d.arg
		}
	case "MinLength":
		applyLengthConstraint(d.arg, &f.minLength)
	case "MaxLength":
		applyLengthConstraint(d.arg, &f.maxLength)
	}
}

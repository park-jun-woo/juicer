//ff:func feature=scan type=extract control=iteration dimension=1 topic=nestjs
//ff:what @ApiProperty({ enum: [...] }) 데코레이터에서 enum 리터럴 배열을 추출한다
package nestjs

import sitter "github.com/smacker/go-tree-sitter"

// extractEnumFromApiProperty extracts literal enum values from @ApiProperty({ enum: ['a', 'b'] }).
// Only handles literal string/number arrays; does not resolve enum type references.
func extractEnumFromApiProperty(prop *sitter.Node, src []byte) []string {
	decNodes := childrenOfType(prop, "decorator")
	for _, dn := range decNodes {
		d := parseDecorator(dn, src)
		if d.name != "ApiProperty" {
			continue
		}
		return extractEnumFromDecoratorNode(dn, src)
	}
	return nil
}

//ff:func feature=scan type=extract control=iteration dimension=1 topic=nestjs
//ff:what enum_body에서 멤버 값을 수집한다
package nestjs

import sitter "github.com/smacker/go-tree-sitter"

// collectEnumMemberValues collects member values from an enum_body node.
// For enum_assignment children (value-bearing members like OPEN = 'open'), the string value is extracted.
// For property_identifier children (valueless members like Up, Down), the key name is used.
func collectEnumMemberValues(body *sitter.Node, src []byte) []string {
	var values []string
	for i := 0; i < int(body.ChildCount()); i++ {
		v, ok := extractEnumAssignmentValue(body.Child(i), src)
		if ok {
			values = append(values, v)
		}
	}
	return values
}

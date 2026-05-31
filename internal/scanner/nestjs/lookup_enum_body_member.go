//ff:func feature=scan type=extract control=iteration dimension=1 topic=nestjs
//ff:what enum_body에서 멤버명 키에 해당하는 값을 찾는다
package nestjs

import sitter "github.com/smacker/go-tree-sitter"

// lookupEnumBodyMember scans an enum_body for the member whose key equals
// memberName and returns its value. Returns ("", false) when absent.
func lookupEnumBodyMember(body *sitter.Node, src []byte, memberName string) (string, bool) {
	for i := 0; i < int(body.ChildCount()); i++ {
		child := body.Child(i)
		if enumMemberKeyName(child, src) != memberName {
			continue
		}
		return extractEnumAssignmentValue(child, src)
	}
	return "", false
}

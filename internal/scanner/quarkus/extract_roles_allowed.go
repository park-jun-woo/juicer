//ff:func feature=scan type=extract control=sequence topic=quarkus
//ff:what @RolesAllowed에서 역할을 추출한다
package quarkus

import sitter "github.com/smacker/go-tree-sitter"

func extractRolesAllowed(node *sitter.Node, src []byte) []string {
	ann := findAnnotation(node, src, AnnRolesAllowed)
	if ann == nil {
		return nil
	}
	args := annotationArgs(ann, src)
	if args == nil {
		return nil
	}
	return extractRoleStrings(args, src)
}

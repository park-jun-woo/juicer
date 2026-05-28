//ff:func feature=scan type=extract control=sequence topic=quarkus
//ff:what @PathParam 파라미터를 분류한다
package quarkus

import (
	sitter "github.com/smacker/go-tree-sitter"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func classifyPathParam(param *sitter.Node, src []byte, ep *endpointInfo, typeName, paramName string) {
	ann := findAnnotation(param, src, AnnPathParam)
	name := resolveAnnotationName(ann, src, paramName)
	ep.params = append(ep.params, scanner.Param{
		Name: name,
		Type: javaTypeToOpenAPIString(typeName),
	})
}

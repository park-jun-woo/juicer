//ff:func feature=scan type=extract control=sequence topic=quarkus
//ff:what @FormParam 파라미터를 분류한다
package quarkus

import (
	sitter "github.com/smacker/go-tree-sitter"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func classifyFormParam(param *sitter.Node, src []byte, ep *endpointInfo, typeName, paramName string) {
	ann := findAnnotation(param, src, AnnFormParam)
	name := resolveAnnotationName(ann, src, paramName)
	ep.formParams = append(ep.formParams, scanner.Param{
		Name: name,
		Type: javaTypeToOpenAPIString(typeName),
	})
}

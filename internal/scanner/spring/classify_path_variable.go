//ff:func feature=scan type=extract control=sequence topic=spring
//ff:what @PathVariable 파라미터를 분류한다
package spring

import (
	sitter "github.com/smacker/go-tree-sitter"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func classifyPathVariable(param *sitter.Node, src []byte, ep *endpointInfo, typeName, paramName string) {
	ann := findAnnotation(param, src, AnnPathVariable)
	name := resolveAnnotationName(ann, src, paramName)
	ep.params = append(ep.params, scanner.Param{
		Name: name,
		Type: javaTypeToOpenAPIString(typeName),
	})
}

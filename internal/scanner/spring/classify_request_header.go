//ff:func feature=scan type=extract control=sequence topic=spring
//ff:what @RequestHeader 파라미터를 분류한다
package spring

import (
	sitter "github.com/smacker/go-tree-sitter"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func classifyRequestHeader(param *sitter.Node, src []byte, ep *endpointInfo, typeName, paramName string) {
	ann := findAnnotation(param, src, AnnRequestHeader)
	name := resolveAnnotationName(ann, src, paramName)
	ep.headers = append(ep.headers, scanner.Param{
		Name: name,
		Type: javaTypeToOpenAPIString(typeName),
	})
}

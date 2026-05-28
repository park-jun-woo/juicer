//ff:func feature=scan type=extract control=sequence topic=quarkus
//ff:what @QueryParam 파라미터를 분류한다
package quarkus

import (
	sitter "github.com/smacker/go-tree-sitter"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func classifyQueryParam(param *sitter.Node, src []byte, ep *endpointInfo, typeName, paramName string) {
	ann := findAnnotation(param, src, AnnQueryParam)
	name := resolveAnnotationName(ann, src, paramName)
	defaultVal := extractDefaultValue(param, src)
	p := scanner.Param{
		Name: name,
		Type: javaTypeToOpenAPIString(typeName),
	}
	if defaultVal != "" {
		p.Default = defaultVal
	}
	ep.query = append(ep.query, p)
}

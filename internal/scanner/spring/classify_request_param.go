//ff:func feature=scan type=extract control=sequence topic=spring
//ff:what @RequestParam 파라미터를 분류한다
package spring

import (
	"strings"

	sitter "github.com/smacker/go-tree-sitter"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func classifyRequestParam(param *sitter.Node, src []byte, ep *endpointInfo, typeName, paramName string, imports map[string]string, referrerPath, projectRoot string) {
	ann := findAnnotation(param, src, AnnRequestParam)
	name := resolveAnnotationName(ann, src, paramName)
	defaultVal := ""
	if ann != nil {
		defaultVal = annotationElementValue(ann, src, "defaultValue")
	}
	if defaultVal != "" && strings.Contains(defaultVal, ".") {
		defaultVal = resolveConstantValue(defaultVal, imports, referrerPath, projectRoot)
	}
	p := scanner.Param{
		Name: name,
		Type: javaTypeToOpenAPIString(typeName),
	}
	if defaultVal != "" {
		p.Default = defaultVal
	}
	ep.query = append(ep.query, p)
}

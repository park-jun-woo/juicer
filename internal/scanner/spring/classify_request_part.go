//ff:func feature=scan type=extract control=sequence topic=spring
//ff:what @RequestPart 파라미터를 분류한다
package spring

import (
	sitter "github.com/smacker/go-tree-sitter"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func classifyRequestPart(param *sitter.Node, src []byte, ep *endpointInfo, paramName string) {
	ann := findAnnotation(param, src, AnnRequestPart)
	name := resolveAnnotationName(ann, src, paramName)
	ep.files = append(ep.files, scanner.Param{
		Name: name,
		Type: "string:binary",
	})
}

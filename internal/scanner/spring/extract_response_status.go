//ff:func feature=scan type=extract control=sequence topic=spring
//ff:what @ResponseStatus에서 HTTP 상태 코드를 추출한다
package spring

import sitter "github.com/smacker/go-tree-sitter"

func extractResponseStatus(m *sitter.Node, src []byte, ep *endpointInfo) {
	ann := findAnnotation(m, src, AnnResponseStatus)
	if ann == nil {
		return
	}
	val := firstStringArg(ann, src)
	if val == "" {
		val = annotationElementValue(ann, src, "value")
		if val == "" {
			val = annotationElementValue(ann, src, "code")
		}
	}
	if val == "" {
		args := annotationArgs(ann, src)
		if args != nil {
			val = extractStatusFromArgs(args, src)
		}
	}
	if code, ok := httpStatusAnnotations[val]; ok {
		ep.statusCode = code
	} else if val != "" {
		ep.statusCode = val
	}
}

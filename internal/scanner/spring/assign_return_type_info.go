//ff:func feature=scan type=extract control=sequence topic=spring
//ff:what 반환 타입 정보를 응답에 할당한다
package spring

import "github.com/park-jun-woo/codistill/internal/scanner"

func assignReturnTypeInfo(ep endpointInfo, resp *scanner.Response) {
	oa := javaTypeToOpenAPI(ep.returnType)
	if oa.Type == "" {
		return
	}
	resp.TypeName = ep.returnType
	if ep.returnIsArray {
		resp.Body = "array"
	}
}

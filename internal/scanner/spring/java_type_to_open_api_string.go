//ff:func feature=scan type=convert control=sequence topic=spring
//ff:what Java 타입을 OpenAPI 타입:포맷 문자열로 변환한다
package spring

func javaTypeToOpenAPIString(jtype string) string {
	oa := javaTypeToOpenAPI(jtype)
	if oa.Format != "" {
		return oa.Type + ":" + oa.Format
	}
	return oa.Type
}

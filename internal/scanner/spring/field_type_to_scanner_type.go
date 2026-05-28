//ff:func feature=scan type=convert control=sequence topic=spring
//ff:what Java 필드 타입을 scanner 내부 타입 문자열로 변환한다
package spring

func fieldTypeToScannerType(jtype string) string {
	oa := javaTypeToOpenAPI(jtype)
	if oa.Format != "" && oa.Type == "string" {
		return oa.Type + ":" + oa.Format
	}
	if oa.Type == "array" && oa.Items != "" {
		inner := javaTypeToOpenAPI(oa.Items)
		if inner.Type != "" && inner.Type != "object" {
			return "array:" + inner.Type
		}
		return "[]" + oa.Items
	}
	if oa.Type != "" {
		return oa.Type
	}
	return jtype
}

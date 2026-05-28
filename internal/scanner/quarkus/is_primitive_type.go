//ff:func feature=scan type=extract control=sequence topic=quarkus
//ff:what 타입이 원시 타입(object 아닌 OpenAPI 타입)인지 확인한다
package quarkus

func isPrimitiveType(t string) bool {
	oa := javaTypeToOpenAPI(t)
	return oa.Type != "" && oa.Type != "object"
}

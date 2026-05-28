//ff:func feature=scan type=extract control=sequence topic=dotnet
//ff:what 타입이 원시 타입(object 아닌 OpenAPI 타입)인지 확인한다
package dotnet

func isPrimitiveType(t string) bool {
	oa := csharpTypeToOpenAPI(t)
	return oa.Type != "" && oa.Type != "object"
}

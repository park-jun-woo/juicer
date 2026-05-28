//ff:func feature=scan type=convert control=sequence topic=dotnet
//ff:what C# 타입을 OpenAPI type 문자열로 변환한다
package dotnet

func csharpTypeToOpenAPIType(ctype string) string {
	oa := csharpTypeToOpenAPI(ctype)
	if oa.Type == "" {
		return "string"
	}
	return oa.Type
}

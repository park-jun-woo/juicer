//ff:type feature=scan type=model topic=dotnet
//ff:what 메서드 본문 return 식에서 추출한 응답 정보
package dotnet

type bodyResponse struct {
	status  string
	typeName string
	isArray bool
	found   bool
}

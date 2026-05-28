//ff:type feature=scan type=model topic=dotnet
//ff:what DTO 해석 요청 구조체
package dotnet

type dtoRequest struct {
	typeName    string
	usings      []string
	referrer    string
	projectRoot string
	epIdx       int
	isBody      bool
	isForm      bool
}

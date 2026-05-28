//ff:type feature=scan type=model topic=quarkus
//ff:what DTO 해석 요청 구조체
package quarkus

type dtoRequest struct {
	typeName    string
	imports     map[string]string
	referrer    string
	projectRoot string
	epIdx       int
	isBody      bool
	isQuery     bool
	isForm      bool
}

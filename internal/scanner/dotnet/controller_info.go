//ff:type feature=scan type=model topic=dotnet
//ff:what 컨트롤러 추출 중간 결과 구조체
package dotnet

type controllerInfo struct {
	prefix    string
	className string
	file      string
	absFile   string
	roles     []string
	endpoints []endpointInfo
	usings    []string
}

//ff:type feature=scan type=model topic=spring
//ff:what 컨트롤러 추출 중간 결과 구조체
package spring

type controllerInfo struct {
	prefix     string
	className  string
	file       string
	absFile    string
	roles      []string
	endpoints  []endpointInfo
	imports    map[string]string
	interfaces []string
}

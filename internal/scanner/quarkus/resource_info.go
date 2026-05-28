//ff:type feature=scan type=model topic=quarkus
//ff:what 리소스 클래스 추출 중간 결과 구조체
package quarkus

type resourceInfo struct {
	prefix    string
	className string
	file      string
	absFile   string
	roles     []string
	endpoints []endpointInfo
	imports   map[string]string
}

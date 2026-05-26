//ff:func feature=scan type=extract control=iteration dimension=1
//ff:what ScanResult에서 등록된 엔드포인트 집합을 구축한다
package scanner

import "strings"

func buildRegisteredSet(scanResult *ScanResult) map[string]bool {
	registered := map[string]bool{}
	for _, ep := range scanResult.Endpoints {
		oaPath := ginPathToOpenAPI(ep.Path)
		method := strings.ToLower(ep.Method)
		for _, m := range expandAnyMethod(method) {
			registered[m+"\t"+oaPath] = true
		}
	}
	return registered
}

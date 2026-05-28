//ff:func feature=scan type=extract control=iteration dimension=1 topic=dotnet
//ff:what 단일 파일에서 Minimal API 호출을 추출한다
package dotnet

import "github.com/park-jun-woo/codistill/internal/scanner"

func extractMinimalAPIsFromFile(fi *fileInfo, groups map[string]string) []scanner.Endpoint {
	var endpoints []scanner.Endpoint
	invocations := findAllByType(fi.root, "invocation_expression")
	for _, inv := range invocations {
		ep, ok := matchMapMethod(inv, fi, groups)
		if ok {
			endpoints = append(endpoints, ep)
		}
	}
	return endpoints
}

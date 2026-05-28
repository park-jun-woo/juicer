//ff:func feature=scan type=extract control=iteration dimension=1 topic=dotnet
//ff:what app.MapGet(), app.MapPost() 등 Minimal API 호출을 추출한다
package dotnet

import "github.com/park-jun-woo/codistill/internal/scanner"

func extractMinimalAPIs(files []*fileInfo, groups map[string]string) []scanner.Endpoint {
	var endpoints []scanner.Endpoint
	for _, fi := range files {
		eps := extractMinimalAPIsFromFile(fi, groups)
		endpoints = append(endpoints, eps...)
	}
	return endpoints
}

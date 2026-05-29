//ff:func feature=scan type=extract control=iteration dimension=1 topic=laravel
//ff:what 라우트 파일(routes/api.php, routes/web.php)에서 라우트를 모은다
package laravel

import "strings"

// collectAllRoutes gathers routes from route files (routes/api.php, routes/web.php).
func collectAllRoutes(parsedFiles map[string]*fileInfo) []routeInfo {
	var routes []routeInfo
	routeFiles := []string{
		"routes/api.php",
		"routes/web.php",
	}
	for _, rf := range routeFiles {
		fi, ok := parsedFiles[rf]
		if !ok {
			continue
		}
		prefix := ""
		if strings.HasSuffix(rf, "api.php") {
			prefix = "api"
		}
		routes = append(routes, collectRoutes(*fi, prefix, nil)...)
		routes = append(routes, collectAPIResource(*fi, prefix, nil)...)
		routes = append(routes, extractRouteGroups(*fi, prefix, nil)...)
	}
	return routes
}

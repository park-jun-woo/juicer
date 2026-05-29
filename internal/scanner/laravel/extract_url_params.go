//ff:func feature=scan type=extract control=iteration dimension=1 topic=laravel
//ff:what URL 경로에서 {param} 변수를 추출한다
package laravel

import (
	"regexp"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

var laravelParamRe = regexp.MustCompile(`\{([a-zA-Z_][a-zA-Z0-9_]*)\}`)

// extractURLParams extracts path parameters from a Laravel URL path.
// Laravel already uses {param} format which matches OpenAPI.
func extractURLParams(path string) []scanner.Param {
	matches := laravelParamRe.FindAllStringSubmatch(path, -1)
	if len(matches) == 0 {
		return nil
	}
	var params []scanner.Param
	for _, m := range matches {
		params = append(params, scanner.Param{
			Name: m[1],
			Type: "string",
		})
	}
	return params
}

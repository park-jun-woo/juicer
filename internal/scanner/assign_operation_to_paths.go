//ff:func feature=scan type=extract control=iteration dimension=1
//ff:what 메서드(any 확장 포함)별로 operation을 paths에 할당하고 충돌 시 경고한다
package scanner

import (
	"fmt"
	"os"
	"strings"
)

func assignOperationToPaths(paths map[string]map[string]any, oaPath string, ep Endpoint, op map[string]any) {
	method := strings.ToLower(ep.Method)
	for _, m := range expandAnyMethod(method) {
		if _, dup := paths[oaPath][m]; dup {
			fmt.Fprintf(os.Stderr, "warning: duplicate operation %s %s (handler %q at %s:%d) overwrites a previous one — check route prefix composition\n",
				strings.ToUpper(m), oaPath, ep.Handler, ep.File, ep.Line)
		}
		paths[oaPath][m] = op
	}
}
